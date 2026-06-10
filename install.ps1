# install.ps1 — install the iicu CLI on Windows.
#
#   irm https://raw.githubusercontent.com/elricho/iicu/main/install.ps1 | iex
#
# Environment overrides:
#   IICU_VERSION       install a specific tag (default: latest release)
#   IICU_INSTALL_DIR   install location  (default: %LOCALAPPDATA%\Programs\iicu)
$ErrorActionPreference = 'Stop'

$Repo   = 'elricho/iicu'
$Binary = 'iicu'
$InstallDir = if ($env:IICU_INSTALL_DIR) { $env:IICU_INSTALL_DIR }
              else { Join-Path $env:LOCALAPPDATA 'Programs\iicu' }

# --- detect architecture ---------------------------------------------------
$arch = switch ($env:PROCESSOR_ARCHITECTURE) {
  'AMD64' { 'amd64' }
  'ARM64' { 'arm64' }
  default { throw "Unsupported architecture: $env:PROCESSOR_ARCHITECTURE" }
}

# --- resolve version -------------------------------------------------------
$version = $env:IICU_VERSION
if (-not $version) {
  # Read the /releases/latest redirect to find the newest tag (no API limit).
  $resp = Invoke-WebRequest -UseBasicParsing -MaximumRedirection 0 `
    -Uri "https://github.com/$Repo/releases/latest" -ErrorAction SilentlyContinue
  $loc = $resp.Headers['Location']
  if ($loc -is [array]) { $loc = $loc[0] }
  $version = ($loc -split '/tag/')[-1]
}
if (-not $version) { throw 'Could not determine latest version' }
$num = $version.TrimStart('v')

# --- download & verify -----------------------------------------------------
$archive = "${Binary}_${num}_windows_${arch}.zip"
$base    = "https://github.com/$Repo/releases/download/$version"
$tmp     = Join-Path $env:TEMP "iicu-install-$(Get-Random)"
New-Item -ItemType Directory -Path $tmp -Force | Out-Null
try {
  Write-Host "Downloading $archive ..."
  Invoke-WebRequest -UseBasicParsing -Uri "$base/$archive"      -OutFile "$tmp\$archive"
  Invoke-WebRequest -UseBasicParsing -Uri "$base/checksums.txt" -OutFile "$tmp\checksums.txt"

  Write-Host "Verifying checksum ..."
  $entry = Select-String -Path "$tmp\checksums.txt" -Pattern ([regex]::Escape($archive))
  if (-not $entry) { throw "No checksum entry for $archive" }
  $expected = ($entry.Line -split '\s+')[0].ToLower()
  $actual   = (Get-FileHash "$tmp\$archive" -Algorithm SHA256).Hash.ToLower()
  if ($expected -ne $actual) { throw "Checksum verification failed for $archive" }

  Expand-Archive -Path "$tmp\$archive" -DestinationPath $tmp -Force

  New-Item -ItemType Directory -Path $InstallDir -Force | Out-Null
  Copy-Item "$tmp\$Binary.exe" (Join-Path $InstallDir "$Binary.exe") -Force
}
finally {
  Remove-Item -Recurse -Force $tmp -ErrorAction SilentlyContinue
}

# --- add to PATH (user scope) ----------------------------------------------
$userPath = [Environment]::GetEnvironmentVariable('Path', 'User')
if ($userPath -notlike "*$InstallDir*") {
  [Environment]::SetEnvironmentVariable('Path', "$userPath;$InstallDir", 'User')
  $env:Path += ";$InstallDir"
  Write-Host "Added $InstallDir to your user PATH (restart your shell to pick it up)."
}

Write-Host "Installed $Binary $version to $InstallDir\$Binary.exe"
& (Join-Path $InstallDir "$Binary.exe") --version
