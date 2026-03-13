# Contributing

## General rules

- Keep the package simple and explicit.
- Do not add validation unless there is a very strong reason.
- Do not try to unify request and response models between different API methods.
- Prefer readable code over clever abstractions.

## Workflow

1. Open an issue if you plan a non-trivial change.
2. Keep pull requests focused.
3. Update documentation when public API changes.
4. Add or update examples when behavior changes.

## Code style

- Follow standard Go formatting.
- Keep method behavior direct and predictable.
- Prefer adding a new struct for a method instead of reusing a similar one from another method.

## Verification

Run:

```bash
env TMPDIR=/tmp GOCACHE=/tmp/go-build CGO_ENABLED=0 go test ./...
```

If you cannot run checks in your environment, mention that in the pull request.
