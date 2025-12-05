# Build script for Flutter Takeoff
# Sets version information at build time

param(
    [string]$OutputName = "flutter-installer"
)

# Get version info
$VERSION = "1.0.0"
$BUILD_DATE = Get-Date -Format "yyyy-MM-ddTHH:mm:ssZ"
$GIT_COMMIT = if (Get-Command git -ErrorAction SilentlyContinue) {
    git rev-parse --short HEAD 2>$null
    if ($LASTEXITCODE -ne 0) { "unknown" }
} else { "unknown" }
$GIT_BRANCH = if (Get-Command git -ErrorAction SilentlyContinue) {
    git rev-parse --abbrev-ref HEAD 2>$null
    if ($LASTEXITCODE -ne 0) { "unknown" }
} else { "unknown" }

Write-Host "Building Flutter Takeoff v$VERSION" -ForegroundColor Cyan
Write-Host "  Build Date: $BUILD_DATE" -ForegroundColor Gray
Write-Host "  Git Commit: $GIT_COMMIT" -ForegroundColor Gray
Write-Host "  Git Branch: $GIT_BRANCH" -ForegroundColor Gray
Write-Host ""

# Build with version information
$LDFLAGS = "-X 'flutter_takeoff/pkg/version.BuildDate=$BUILD_DATE' " +
           "-X 'flutter_takeoff/pkg/version.GitCommit=$GIT_COMMIT' " +
           "-X 'flutter_takeoff/pkg/version.GitBranch=$GIT_BRANCH'"

go build -ldflags $LDFLAGS -o "$OutputName.exe" .

if ($LASTEXITCODE -eq 0) {
    Write-Host "Build successful: $OutputName.exe" -ForegroundColor Green
    Write-Host ""
    Write-Host "To run: .\$OutputName.exe" -ForegroundColor Yellow
} else {
    Write-Host "Build failed" -ForegroundColor Red
    exit 1
}
