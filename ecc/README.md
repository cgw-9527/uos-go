UOSIO Elliptic Curve Cryptography Wrapper
=========================================

This is a simple wrapper for `btcec`, that handles the specificities
of the format for keys in UOS.

It was crafted in reference to `uosjs-ecc`, `uosjs-keygen` and the C++
codebase of UOS.IO Software.

This handles the `UOS` prefix on public keys, manages the version and
checksums in public and private keys.
