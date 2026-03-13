# Changelog

## [Unreleased]

### Changed
- Added `Wait()` call to wait for Redis to be ready before returning client
- Added port exposure in container config
- Added `RedisUpWaitTime` variable (default: 10 seconds) to configure wait timeout
- Updated to use Go modules
