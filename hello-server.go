name: Go Build and Test

on:
  push:
    branches:
      - main
  schedule:
    - cron: '*/15 * * * 1-6' # Executa a cada 15 minutos de segunda a sábado

jobs:
  build:
    runs-on: ubuntu-latest
    env:
      FILE_NAME: hello-server

    steps:
      - name: Checkout code
        uses: actions/checkout@v3

      - name: Build Go artifact for Linux
        run: go build ${{ env.FILE_NAME }}.go

      - name: Build Go artifact for Windows
        run: GOOS=windows GOARCH=amd64 go build ${{ env.FILE_NAME }}.go

      - name: Upload Linux artifact
        uses: actions/upload-artifact@v3
        with:
          name: linux-artifact
          path: ${{ env.FILE_NAME }}

      - name: Upload Windows artifact
        uses: actions/upload-artifact@v3
        with:
          name: windows-artifact
          path: ${{ env.FILE_NAME }}.exe
