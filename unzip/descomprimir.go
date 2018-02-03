package main

import (
    "archive/zip"
    "fmt"
    "io"
    "log"
    "os"
    "path/filepath"
    "strings"
)



func main() {

    files, err := Unzip("Archivo.zip", "output")

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Unzipped: " + strings.Join(files, ", "))

}







//un-compress
// moving all files and folders to a directory
func Unzip(src, dest string) ([]string, error) {

    var filenames []string

    r, err := zip.OpenReader(src)

    //si hay error en la descompresión
    if err != nil {
        return filenames, err
    }
    defer r.Close()


    for _, f := range r.File {

        rc, err := f.Open()
        if err != nil {
            return filenames, err
        }
        defer rc.Close()

        // Store filename/path for returning and using later on
        fpath := filepath.Join(dest, f.Name)
        filenames = append(filenames, fpath)

        if f.FileInfo().IsDir() {

            // Crea el folder
            os.MkdirAll(fpath, os.ModePerm)

        } else {

            // Make File
            var fdir string
            if lastIndex := strings.LastIndex(fpath, string(os.PathSeparator)); lastIndex > -1 {
                fdir = fpath[:lastIndex]
            }

            err = os.MkdirAll(fdir, os.ModePerm)
            if err != nil {
                log.Fatal(err)
                return filenames, err
            }
            f, err := os.OpenFile(
                fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
            if err != nil {
                return filenames, err
            }
            defer f.Close()

            _, err = io.Copy(f, rc)
            if err != nil {
                return filenames, err
            }

        }
    }
    return filenames, nil
}