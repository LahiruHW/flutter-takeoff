# Flutter Installation Guide

This document provides detailed information about installing Flutter and its dependencies manually, which complements what Flutter Takeoff automates.

## Prerequisites for Windows + Android Development

### 1. Git for Windows

**Why needed**: Flutter SDK is distributed via Git, and Flutter's tools use Git for version management.

**Installation**:
1. Download from: https://git-scm.com/download/win
2. Run the installer
3. Use recommended settings (or customize as needed)
4. Verify installation:
   ```powershell
   git --version
   ```

**Expected output**: `git version 2.x.x`

### 2. Java Development Kit (JDK)

**Why needed**: Required for building Android apps and running Android command-line tools.

**Recommended**: OpenJDK 17 or later (LTS version)

**Installation**:
1. Download Eclipse Temurin from: https://adoptium.net/
2. Select Java 17 (LTS) or Java 21 (LTS)
3. Download and run the Windows installer (.msi)
4. During installation, select "Add to PATH"
5. Verify installation:
   ```powershell
   java -version
   ```

**Expected output**: `openjdk version "17.x.x"` or similar

**Set JAVA_HOME** (if not set automatically):
```powershell
setx JAVA_HOME "C:\Program Files\Eclipse Adoptium\jdk-17.x.x-hotspot"
```

### 3. Android SDK

**Why needed**: Provides tools for building, testing, and deploying Android apps.

#### Option A: Android Studio (Recommended for Beginners)

1. Download from: https://developer.android.com/studio
2. Run the installer
3. During setup, ensure you install:
   - Android SDK
   - Android SDK Platform
   - Android Virtual Device (AVD)

4. After installation, open Android Studio
5. Go to: Tools → SDK Manager
6. Install:
   - Android SDK Platform 36 (or latest)
   - Android SDK Build-Tools
   - Android SDK Command-line Tools
   - Android SDK Platform-Tools
   - Android Emulator

#### Option B: Command-line Tools Only

1. Download command-line tools from: https://developer.android.com/studio#command-tools
2. Extract to a location like `C:\Android\cmdline-tools`
3. Set environment variables:
   ```powershell
   setx ANDROID_SDK_ROOT "C:\Android\sdk"
   setx PATH "%PATH%;%ANDROID_SDK_ROOT%\platform-tools;%ANDROID_SDK_ROOT%\cmdline-tools\latest\bin"
   ```

4. Install SDK packages:
   ```powershell
   sdkmanager "platform-tools" "platforms;android-34" "build-tools;34.0.0" "cmdline-tools;latest"
   ```

### 4. Flutter SDK

**Installation Steps**:

1. **Download Flutter**:
   - Visit: https://docs.flutter.dev/get-started/install/windows
   - Download the latest stable release ZIP file
   - Or use direct link: https://storage.googleapis.com/flutter_infra_release/releases/stable/windows/flutter_windows_3.x.x-stable.zip

2. **Extract Flutter**:
   - Extract the zip file to a desired location (e.g., `C:\src\flutter`)
   - **Important**: Do NOT install to `C:\Program Files\` (requires elevated privileges)

3. **Update PATH**:
   ```powershell
   setx PATH "%PATH%;C:\src\flutter\bin"
   ```

4. **Restart Terminal**: Close and reopen your terminal/PowerShell

5. **Run Flutter Doctor**:
   ```powershell
   flutter doctor
   ```

6. **Accept Android Licenses**:
   ```powershell
   flutter doctor --android-licenses
   ```
   Press 'y' to accept each license

## Verification Steps

### Check All Dependencies

Run Flutter Doctor with verbose output:
```powershell
flutter doctor -v
```

**Expected output** (all checkmarks):
```
Doctor summary (to see all details, run flutter doctor -v):
[√] Flutter (Channel stable, 3.x.x, on Microsoft Windows...)
[√] Windows Version (Installed version of Windows is version 10...)
[√] Android toolchain - develop for Android devices (Android SDK version 34.x.x)
[√] Chrome - develop for the web
[√] Visual Studio - develop Windows apps (Visual Studio Community 2022...)
[√] Android Studio (version 2024.x)
[√] VS Code (version 1.x.x)
[√] Connected device (1 available)
[√] Network resources
```

### Test Flutter Installation

Create a test project:
```powershell
flutter create test_app
cd test_app
flutter run
```

## Common Issues and Solutions

### Issue: 'flutter' is not recognized

**Solution**: 
1. Verify Flutter is in PATH:
   ```powershell
   echo $env:PATH
   ```
2. Restart terminal
3. If still not working, manually add to PATH again

### Issue: Android licenses not accepted

**Solution**:
```powershell
flutter doctor --android-licenses
```
Accept all licenses by pressing 'y'

### Issue: cmdline-tools not found

**Solution**:
1. Open Android Studio
2. Tools → SDK Manager → SDK Tools
3. Check "Android SDK Command-line Tools (latest)"
4. Click Apply

### Issue: No connected devices

**Solution**:
1. For emulator:
   ```powershell
   # List available emulators
   flutter emulators
   
   # Launch an emulator
   flutter emulators --launch <emulator_id>
   ```

2. For physical device:
   - Enable Developer Options on Android device
   - Enable USB Debugging
   - Connect via USB
   - Verify: `flutter devices`

### Issue: Java version issues

**Solution**:
```powershell
# Check current version
java -version

# Set JAVA_HOME to correct version
setx JAVA_HOME "C:\Program Files\Eclipse Adoptium\jdk-17.x.x-hotspot"
```

## Environment Variables Reference

| Variable | Purpose | Example Value |
|----------|---------|---------------|
| `JAVA_HOME` | Java JDK location | `C:\Program Files\Eclipse Adoptium\jdk-17.0.2-hotspot` |
| `ANDROID_HOME` | Android SDK location | `C:\Users\YourName\AppData\Local\Android\Sdk` |
| `ANDROID_SDK_ROOT` | Android SDK location (newer) | Same as ANDROID_HOME |
| `PATH` | Executable locations | Should include Flutter bin, Android platform-tools |

## Setting Environment Variables

### Temporary (Current Session Only)
```powershell
$env:JAVA_HOME = "C:\Program Files\Eclipse Adoptium\jdk-17.x.x-hotspot"
```

### Permanent (User Level)
```powershell
setx JAVA_HOME "C:\Program Files\Eclipse Adoptium\jdk-17.x.x-hotspot"
```

### Permanent (System Level - Requires Admin)
```powershell
setx JAVA_HOME "C:\Program Files\Eclipse Adoptium\jdk-17.x.x-hotspot" /M
```

### Using GUI
1. Open System Properties (Win + Pause/Break)
2. Click "Advanced system settings"
3. Click "Environment Variables"
4. Add/Edit variables under "User variables" or "System variables"

## Flutter Commands Reference

### Basic Commands
```powershell
# Check Flutter installation
flutter doctor

# Detailed diagnostics
flutter doctor -v

# List available devices
flutter devices

# Create new project
flutter create my_app

# Run app
flutter run

# Build APK
flutter build apk

# Clean build files
flutter clean

# Update Flutter
flutter upgrade
```

### Channel Management
```powershell
# See current channel
flutter channel

# Switch to stable
flutter channel stable

# Switch to beta
flutter channel beta

# Get latest from current channel
flutter upgrade
```

## Creating Your First App

```powershell
# Create app
flutter create hello_flutter
cd hello_flutter

# Run on Chrome (web)
flutter run -d chrome

# Run on Android emulator (start emulator first)
flutter run -d android

# Run on specific device
flutter run -d <device_id>
```

## Android Emulator Setup

### Create Emulator (GUI)
1. Open Android Studio
2. Tools → Device Manager
3. Create Virtual Device
4. Select hardware (e.g., Pixel 5)
5. Download system image (e.g., Android 13)
6. Finish setup

### Create Emulator (CLI)
```powershell
# List available system images
sdkmanager --list | findstr system-images

# Download system image
sdkmanager "system-images;android-33;google_apis;x86_64"

# Create AVD
avdmanager create avd -n flutter_emulator -k "system-images;android-33;google_apis;x86_64"

# List emulators
flutter emulators

# Launch emulator
flutter emulators --launch flutter_emulator
```

## VS Code Setup (Optional but Recommended)

1. Install VS Code: https://code.visualstudio.com/
2. Install Flutter extension:
   - Open VS Code
   - Go to Extensions (Ctrl+Shift+X)
   - Search for "Flutter"
   - Install "Flutter" by Dart Code

3. Configure Flutter:
   - Ctrl+Shift+P → "Flutter: Select Device"
   - Choose your target device

## Performance Tips

### Speed up emulator
1. Enable Hardware Acceleration (HAXM or Hyper-V)
2. Allocate more RAM to emulator
3. Use x86_64 system images (faster than ARM)

### Speed up builds
```powershell
# Enable Gradle daemon
echo "org.gradle.daemon=true" >> android/gradle.properties

# Increase Gradle memory
echo "org.gradle.jvmargs=-Xmx4096m" >> android/gradle.properties
```

## Next Steps

After successful installation:

1. ✓ Follow official Flutter codelabs: https://docs.flutter.dev/codelabs
2. ✓ Read Flutter documentation: https://docs.flutter.dev/
3. ✓ Join Flutter community: https://flutter.dev/community
4. ✓ Explore sample apps: https://flutter.github.io/samples/

## Resources

- Flutter Documentation: https://docs.flutter.dev/
- Flutter YouTube Channel: https://www.youtube.com/@flutterdev
- Flutter Community: https://flutter.dev/community
- Dart Language Tour: https://dart.dev/guides/language/language-tour
- Flutter Samples: https://flutter.github.io/samples/

---

Good luck with your Flutter development journey! 🎯
