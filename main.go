package main

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	checkAH94Installed()
	checkEF24GInstalled()
	removeAH94Files()
	removeEF24GFiles()
	unpackFiles()
	install()
	if err := zstd("-d --patch-from=" + vtolvrpath + "\\VTOLVR_Data\\resources.resource " + vtolvrpath + "\\VTOLVR_Data\\resources.resource.patch -o " + vtolvrpath + "\\VTOLVR_Data\\resources.resource.mod"); err != nil {
		log.Fatal(err)
	}

	if err := zstd("-d --patch-from=" + vtolvrpath + "\\VTOLVR_Data\\resources.assets.resS " + vtolvrpath + "\\VTOLVR_Data\\resources.assets.resS.patch -o " + vtolvrpath + "\\VTOLVR_Data\\resources.assets.resS.mod"); err != nil {
		log.Fatal(err)
	}

	if err := zstd("-d --patch-from=" + vtolvrpath + "\\VTOLVR_Data\\resources.assets " + vtolvrpath + "\\VTOLVR_Data\\resources.assets.patch -o " + vtolvrpath + "\\VTOLVR_Data\\resources.assets.mod"); err != nil {
		log.Fatal(err)
	}

	cleanup()
}

func zstd(arguments string) error {
	// Define the PowerShell command to decompress the file using zstd.exe
	cmd := exec.Command("powershell", "-Command", fmt.Sprintf(`& {./zstd.exe "%s"}`, arguments))

	// Run the command and capture any errors
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to decompress file: %v\n%s", err, output)
	}

	fmt.Println("File decompressed successfully.")
	return nil
}

func cleanup() {
	os.Remove(".\\resources.assets.patch")
	os.Remove(".\\resources.assets.resS.patch")
	os.Remove(".\\resources.resource.patch")
	os.Remove(".\\1770480")
	os.Remove(".\\2531290")
	os.Remove(".\\zstd.exe")
}

func unpackFiles() {
	if err := os.WriteFile(".\\installer.zip", installerfiles, 0644); err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	err := unzip(".\\installer.zip", ".\\")

	if err != nil {
		log.Fatal(err)
	}

	ah94, err = os.ReadFile(".\\1770480")
	if err != nil {
		log.Fatal(err)
	}

	ef24g, err = os.ReadFile(".\\2531290")
	if err != nil {
		log.Fatal(err)
	}

	resourcesassetspatch, err = os.ReadFile(".\\resources.assets.patch")
	if err != nil {
		log.Fatal(err)
	}
	resourcesassetsresspatch, err = os.ReadFile(".\\resources.assets.resS.patch")
	if err != nil {
		log.Fatal(err)
	}
	resourcesresourcepatch, err = os.ReadFile(".\\resources.resource.patch")
	if err != nil {
		log.Fatal(err)
	}

}

func install() {
	var err error
	vtolvrpath, err = readLibraryPaths()
	if err != nil {
		log.Fatal(err)
	}
	vtolvrpath += "\\steamapps\\common\\VTOL VR"

	if err := os.WriteFile(vtolvrpath+"\\VTOLVR_Data\\resources.resource.patch", resourcesresourcepatch, 0644); err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
	if err := os.WriteFile(vtolvrpath+"\\VTOLVR_Data\\resources.assets.patch", resourcesassetspatch, 0644); err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
	if err := os.WriteFile(vtolvrpath+"\\VTOLVR_Data\\resources.assets.resS.patch", resourcesassetsresspatch, 0644); err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
	if AH94Installed {
		if err := os.WriteFile(vtolvrpath+"\\DLC\\1770480\\1770480", ah94, 0644); err != nil {
			fmt.Println("Error writing file:", err)
			return
		}

	}
	if EF24GInstalled {
		if err := os.WriteFile(vtolvrpath+"\\DLC\\2531290\\2531290", ef24g, 0644); err != nil {
			fmt.Println("Error writing file:", err)
			return
		}
	}
}

func removeAH94Files() {
	libraryPath, err := readLibraryPaths()
	if err != nil {
		log.Fatal(err)
	}
	err = os.Remove(libraryPath + "\\steamapps\\common\\VTOL VR\\DLC\\1770480\\1770480")
	if err != nil {
		fmt.Println(err)
	}
}

func removeEF24GFiles() {
	libraryPath, err := readLibraryPaths()
	if err != nil {
		log.Fatal(err)
	}
	err = os.Remove(libraryPath + "\\steamapps\\common\\VTOL VR\\DLC\\2531290\\2531290")
	if err != nil {
		fmt.Println(err)
	}
}

func checkAH94Installed() {
	libraryPath, err := readLibraryPaths()
	if err != nil {
		log.Fatal(err)
	}
	exist := exists(libraryPath + "\\steamapps\\common\\VTOL VR\\DLC\\1770480")
	if !exist {
		AH94Installed = false
	} else {
		AH94Installed = true
	}
	exist = exists(libraryPath + "\\steamapps\\common\\VTOL VR\\DLC\\1770480\\1770480.manifest")
	if !exist {
		AH94Installed = false
	} else {
		AH94Installed = true
	}
}

func checkEF24GInstalled() {
	libraryPath, err := readLibraryPaths()
	if err != nil {
		log.Fatal(err)
	}
	exist := exists(libraryPath + "\\steamapps\\common\\VTOL VR\\DLC\\2531290")
	if !exist {
		EF24GInstalled = false
	} else {
		EF24GInstalled = true
	}
	exist = exists(libraryPath + "\\steamapps\\common\\VTOL VR\\DLC\\2531290\\2531290.manifest")
	if !exist {
		EF24GInstalled = false
	} else {
		EF24GInstalled = true
	}
}
