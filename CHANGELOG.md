# Changelog

## [Released] v0.0.1

### Initial Commit
- **Added**: Basic structure of the HDIUtiler application.
- **Implemented**:
  - **Integrity Check**: Verifies the DMG file integrity using `hdiutil verify`.
  - **Mounting**: Mounts the DMG file to a predefined mount point (`/Volumes/RandomMountPointName`).
  - **Application Location**: Locates the application within the mounted DMG.
  - **Code Signing Verification**: Checks if the application is properly signed and trusted.
  - **Unmounting**: Ensures the DMG is unmounted after processing.
- **Dependencies**:
  - `hdiutil` for handling disk images.
  - `dmgutils` package for DMG-related utilities.
  - `utils` package for helper functions.
- **Error Handling**: Includes error messages and handling for command execution, DMG mounting, application location, and code signing verification.

---

*This is the initial commit. Future changes and updates will be recorded here.*
