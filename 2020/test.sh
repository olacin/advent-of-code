find . -maxdepth 1 -type d -name 'day*' -exec sh -c 'cd "{}" && cargo test' \;
