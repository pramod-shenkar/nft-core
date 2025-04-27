package storage

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath" // Import for path manipulation

	"github.com/ipfs/boxo/coreiface/path"
	"github.com/ipfs/go-cid"
	files "github.com/ipfs/go-libipfs/files" // Alias to avoid shadowing
	"github.com/ipfs/kubo/client/rpc"
)

type Client struct {
	*rpc.HttpApi
	tempDir string // Store the temporary directory path
}

func NewClient() (*Client, error) {
	// Use 5001, which is the default IPFS port.  8000 is often used for the gateway.
	httpApi, err := rpc.NewURLApiWithClient("localhost:5001", http.DefaultClient)
	if err != nil {
		return nil, err
	}

	// Create a temporary directory for storing retrieved files.
	tempDir, err := os.MkdirTemp("", "ipfs_storage_")
	if err != nil {
		return nil, fmt.Errorf("failed to create temporary directory: %w", err)
	}

	client := &Client{
		HttpApi: httpApi,
		tempDir: tempDir,
	}

	return client, nil
}

// Close cleans up resources used by the client, such as the temporary directory.
func (c *Client) Close() error {
	if c.tempDir != "" {
		if err := os.RemoveAll(c.tempDir); err != nil {
			return fmt.Errorf("failed to remove temporary directory: %w", err)
		}
		c.tempDir = "" // prevent double removal
	}
	return nil
}

func (c *Client) Get(ctx context.Context, storageId string) ([]byte, error) {
	id, err := cid.Decode(storageId)
	if err != nil {
		return nil, fmt.Errorf("invalid CID: %w", err)
	}

	node, err := c.Unixfs().Get(ctx, path.IpfsPath(id))
	if err != nil {
		return nil, fmt.Errorf("failed to get node from IPFS: %w", err)
	}
	defer func() {
		if closer, ok := node.(io.Closer); ok {
			closer.Close()
		}
	}()

	// Use filepath.Join for cross-platform compatibility
	savedAt := filepath.Join(c.tempDir, storageId)

	err = files.WriteTo(node, savedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to write node to file: %w", err)
	}
	defer os.Remove(savedAt) //Remove temp file

	payload, err := os.ReadFile(savedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	return payload, nil
}

func (c *Client) Save(ctx context.Context, data []byte) (string, error) {
	// files.NewBytesFile takes a byte slice, so this is correct.
	p, err := c.Unixfs().Add(ctx, files.NewBytesFile(data))
	if err != nil {
		return "", fmt.Errorf("failed to add data to IPFS: %w", err)
	}
	return p.Cid().String(), nil
}
