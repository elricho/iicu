#!/usr/bin/env sh
# install.sh — install the iicu CLI on macOS / Linux.
#
#   curl -fsSL https://raw.githubusercontent.com/elricho/iicu/main/install.sh | sh
#
# Environment overrides:
#   IICU_VERSION       install a specific tag (default: latest release)
#   IICU_INSTALL_DIR   install location  (default: /usr/local/bin)
set -eu

REPO="elricho/iicu"
BINARY="iicu"
INSTALL_DIR="${IICU_INSTALL_DIR:-/usr/local/bin}"

err() { echo "install.sh: $*" >&2; exit 1; }

# --- detect platform -------------------------------------------------------
os=$(uname -s | tr '[:upper:]' '[:lower:]')
case "$os" in
  darwin | linux) ;;
  *) err "unsupported OS '$os' — use install.ps1 on Windows" ;;
esac

arch=$(uname -m)
case "$arch" in
  x86_64 | amd64) arch="amd64" ;;
  arm64 | aarch64) arch="arm64" ;;
  *) err "unsupported architecture '$arch'" ;;
esac

# --- resolve version -------------------------------------------------------
version="${IICU_VERSION:-}"
if [ -z "$version" ]; then
  # Follow the /releases/latest redirect to find the newest tag (no API limit).
  version=$(curl -fsSLI -o /dev/null -w '%{url_effective}' \
    "https://github.com/$REPO/releases/latest" | sed 's#.*/tag/##')
fi
[ -n "$version" ] || err "could not determine latest version"
num="${version#v}"

# --- download & verify -----------------------------------------------------
archive="${BINARY}_${num}_${os}_${arch}.tar.gz"
base="https://github.com/$REPO/releases/download/$version"
tmp=$(mktemp -d)
trap 'rm -rf "$tmp"' EXIT

echo "Downloading $archive ..."
curl -fsSL "$base/$archive" -o "$tmp/$archive" || err "download failed: $archive"
curl -fsSL "$base/checksums.txt" -o "$tmp/checksums.txt" || err "download failed: checksums.txt"

echo "Verifying checksum ..."
(
  cd "$tmp"
  line=$(grep " $archive\$" checksums.txt) || err "no checksum entry for $archive"
  echo "$line" | sha256sum -c - >/dev/null 2>&1 ||
    echo "$line" | shasum -a 256 -c - >/dev/null 2>&1 ||
    err "checksum verification failed for $archive"
)

tar -xzf "$tmp/$archive" -C "$tmp"

# --- install ---------------------------------------------------------------
if [ -w "$INSTALL_DIR" ]; then
  mv "$tmp/$BINARY" "$INSTALL_DIR/$BINARY"
else
  echo "Elevated permissions needed to write to $INSTALL_DIR"
  sudo mv "$tmp/$BINARY" "$INSTALL_DIR/$BINARY"
fi
chmod +x "$INSTALL_DIR/$BINARY" 2>/dev/null || true

echo "Installed $BINARY $version to $INSTALL_DIR/$BINARY"
"$INSTALL_DIR/$BINARY" --version || true
