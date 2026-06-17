package main

import (
  "fmt"
  "net"
  "net/http"
  "os"
  "os/exec"
  "path/filepath"
  "runtime"
  "time"
)

func main() {
  if runtime.GOOS != "windows" { fmt.Println("Tento spouštěč je určen pro Windows."); return }
  exe, err := os.Executable(); if err != nil { fmt.Println("Nelze zjistit umístění aplikace:", err); pause(); return }
  root := filepath.Dir(exe)
  dist := filepath.Join(root, "dist")
  if _, err := os.Stat(filepath.Join(dist, "index.html")); err != nil {
    fmt.Println("CHYBA: Chybí produkční soubory ve složce dist.")
    fmt.Println("Rozbalte celý ZIP do jedné složky a spusťte aplikaci znovu.")
    pause(); return
  }

  mux := http.NewServeMux()
  fs := http.FileServer(http.Dir(dist))
  mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    requested := filepath.Join(dist, filepath.Clean(r.URL.Path))
    if r.URL.Path != "/" {
      if info, err := os.Stat(requested); err == nil && !info.IsDir() { fs.ServeHTTP(w, r); return }
    }
    http.ServeFile(w, r, filepath.Join(dist, "index.html"))
  })

  listener, err := net.Listen("tcp", "127.0.0.1:0")
  if err != nil { fmt.Println("CHYBA: Nelze spustit lokální server:", err); pause(); return }
  addr := "http://" + listener.Addr().String()
  fmt.Println("=== KLIK – Kooperativní likvidační informační komplex ===")
  fmt.Println("Aplikace běží na", addr)
  fmt.Println("Toto okno ponechte otevřené. Zavřením aplikaci ukončíte.")

  go func(){
    time.Sleep(500 * time.Millisecond)
    exec.Command("rundll32", "url.dll,FileProtocolHandler", addr).Start()
  }()
  if err := http.Serve(listener, mux); err != nil { fmt.Println("Server byl ukončen:", err) }
}

func pause(){ fmt.Println("Stiskněte Enter pro ukončení."); fmt.Scanln() }
