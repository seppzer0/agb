package manager

import (
	"fmt"
	"io"
	"net/http"
	"os"

	cerror "agb/error"
	"agb/tool"
)

type FileManager struct{}

func NewFileManager() *FileManager {
	return &FileManager{}
}

// Download downloads a file from specified URL
func (fm *FileManager) Download(url string, path string) error {
	resp, err := http.Get(url)
	if err != nil {
		return cerror.ErrGeneric{
			Message: fmt.Sprintf("Could not complete a download from URL (%s)", url),
		}
	}
	defer resp.Body.Close()

	out, err := os.Create(path)
	if err != nil {
		return cerror.ErrGeneric{
			Message: fmt.Sprintf("Could not create a destination file (%s)", path),
		}
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

// Delete removes file or directory from specified path
func (fm *FileManager) Delete(path string) error {
	fileInfo, err_read := os.Stat(path)
	if err_read != nil {
		return cerror.ErrGeneric{
			Message: fmt.Sprintf("There is a problem with path (%s), please check manually", path),
		}
	}

	var err_rm error

	if fileInfo.IsDir() {
		err_rm = os.RemoveAll(path)
	} else {
		err_rm = os.Remove(path)
	}
	if err_rm != nil {
		return cerror.ErrGeneric{
			Message: fmt.Sprintf("Could not delete (%s) path", path),
		}
	}

	return nil
}

// Ucopy universally copies files into desired destination
func (fm *FileManager) Ucopy(path_s string, path_t string) error {
	fileInfo, err_read := os.Stat(path_s)
	if err_read != nil {
		return cerror.ErrGeneric{
			Message: fmt.Sprintf("There is a problem with path (%s), please check manually", path_s),
		}
	}

	if fileInfo.IsDir() {
		fmt.Println("wow")
	} else {
		fmt.Println("wow")
	}

	return nil
}

// UnpackTarGz unpacks a .tar.gz file into specified path
func (fm *FileManager) UnpackTarGz(path_s string, path_t string) error {

	cmd := fmt.Sprintf("tar -xvzf %s -C %s", path_s, path_t)

	out, err := tool.RunCmd(cmd)
	if err != nil {
		return cerror.ErrCommandRun{Command: cmd, Output: out}
	}

	return nil
}
