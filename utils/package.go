package utils

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/viper"
)

type Package struct {
	Name    string
	Version float64

	FullName string

	WinDlLink   string
	MacDlLink   string
	LinuxDlLink string

	Triggers []string

	Installed bool
}

func DownloadAndReadManifest(name string, version float64, repo string, modsFolder string) (*Package, error) {

	file, ferr := os.Open(os.TempDir() + "/midget/packages.csv")
	if ferr != nil {
		return nil, ferr
	}

	// searches packages by name

	sc := bufio.NewScanner(file)

	versions := []float64{}

	for sc.Scan() {
		item := strings.Split(sc.Text(), ",")

		if strings.Contains(item[0], name) {
			ver, err := strconv.ParseFloat(item[1], 64)
			if err != nil {
				return nil, err
			}

			versions = append(versions, ver)
		}
	}

	// sorts versions
	sort.Slice(versions, func(i, j int) bool { return versions[i] > versions[j] })

	if version != 0 {
		var url string
		var path string

		if version == float64(int(version)) {
			fmt.Printf("Found package %v version %v.0\n", name, version)

			fmt.Println("Downloading manifest...")

			url = fmt.Sprintf("https://raw.githubusercontent.com/%v/main/%v/%v.0/manifest.yaml", repo, name, version)
			path = fmt.Sprintf("%v/midget/%v/%v.0/manifest.yaml", os.TempDir(), name, version)
		} else {
			fmt.Printf("Found package %v version %v\n", name, version)

			fmt.Println("Downloading manifest...")

			url = fmt.Sprintf("https://raw.githubusercontent.com/%v/main/%v/%v/manifest.yaml", repo, name, version)
			path = fmt.Sprintf("%v/midget/%v/%v/manifest.yaml", os.TempDir(), name, version)
		}

		err := DownloadFile(url, path)
		if err != nil {
			return nil, err
		}

		pkg, err := ReadManifest(path, modsFolder)
		if err != nil {
			return nil, err
		}

		return pkg, nil
	} else {
		var url string
		var path string

		if versions[0] == float64(int(versions[0])) {
			fmt.Printf("Found package %v version %v.0\n", name, versions[0])

			fmt.Println("Downloading manifest...")

			url = fmt.Sprintf("https://raw.githubusercontent.com/%v/main/%v/%v.0/manifest.yaml", repo, name, versions[0])
			path = fmt.Sprintf("%v/midget/%v/%v.0/manifest.yaml", os.TempDir(), name, versions[0])
		} else {
			fmt.Printf("Found package %v version %v\n", name, versions[0])

			fmt.Println("Downloading manifest...")

			url = fmt.Sprintf("https://raw.githubusercontent.com/%v/main/%v/%v/manifest.yaml", repo, name, versions[0])
			path = fmt.Sprintf("%v/midget/%v/%v/manifest.yaml", os.TempDir(), name, versions[0])
		}

		err := DownloadFile(url, path)
		if err != nil {
			return nil, err
		}

		fmt.Println("Reading manifest...")

		pkg, err := ReadManifest(path, modsFolder)
		if err != nil {
			return nil, err
		}

		return pkg, nil
	}
}

func ReadManifest(path string, modsFolder string) (*Package, error) {
	fmt.Println("Reading manifest...")

	var isInstalled bool

	vi := viper.New()

	vi.SetConfigFile(path)
	vi.SetConfigType("yaml")

	err := vi.ReadInConfig()
	if err != nil {
		return nil, err
	}

	if vi.GetString("win_dl_link") == "" {
		return nil, fmt.Errorf("windows download link not found")
	}

	_, err = os.Stat(modsFolder + "/" + vi.GetString("package"))
	if os.IsNotExist(err) {
		isInstalled = false
	} else {
		isInstalled = true
	}

	return &Package{
		Name:    vi.GetString("package"),
		Version: vi.GetFloat64("version"),

		FullName: vi.GetString("full_name"),

		WinDlLink:   vi.GetString("win_dl_link"),
		MacDlLink:   vi.GetString("mac_dl_link"),
		LinuxDlLink: vi.GetString("linux_dl_link"),

		Triggers: vi.GetStringSlice("triggers"),

		Installed: isInstalled,
	}, nil
}

func (pkg *Package) DownloadPackage() error {
	fmt.Printf("Downloading package %s...\n", pkg.Name)

	var path string

	if pkg.Version == float64(int(pkg.Version)) {
		path = fmt.Sprintf("%v/midget/%v-%v.0.zip", os.TempDir(), pkg.Name, pkg.Version)
	} else {
		path = fmt.Sprintf("%v/midget/%v-%v.zip", os.TempDir(), pkg.Name, pkg.Version)
	}

	switch runtime.GOOS {
	case "windows":
		err := DownloadFile(pkg.WinDlLink, path)
		if err != nil {
			return err
		}
	case "darwin":
		if pkg.MacDlLink == "" {
			LogWarning("Mac version not found, downloading windows version...")
			err := DownloadFile(pkg.WinDlLink, path)
			if err != nil {
				return err
			}
		} else {
			err := DownloadFile(pkg.MacDlLink, path)
			if err != nil {
				return err
			}
		}

	case "linux":
		if pkg.LinuxDlLink == "" {
			LogWarning("Linux version not found, downloading windows version...")
			err := DownloadFile(pkg.WinDlLink, path)
			if err != nil {
				return err
			}
		} else {
			err := DownloadFile(pkg.LinuxDlLink, path)
			if err != nil {
				return err
			}
		}

	default:
		err := DownloadFile(pkg.WinDlLink, path)
		if err != nil {
			return err
		}
	}

	return nil
}

func (pkg *Package) InstallPackage(modsFolder string) error {
	var path string

	if pkg.Version == float64(int(pkg.Version)) {
		path = fmt.Sprintf("%v/midget/%v-%v.0.zip", os.TempDir(), pkg.Name, pkg.Version)
	} else {
		path = fmt.Sprintf("%v/midget/%v-%v.zip", os.TempDir(), pkg.Name, pkg.Version)
	}

	dst := fmt.Sprintf("%v/%v", modsFolder, pkg.Name)

	derr := pkg.DownloadPackage()
	if derr != nil {
		return derr
	}

	err := UnZip(path, dst)
	if err != nil {
		return err
	}

	return nil
}
