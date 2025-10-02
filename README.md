# LRU Cache

A bounded in-memory least-recently-used cache for integer key/value pairs.

## Design

- `map[int]*list.Element` provides `O(1)` access by key.
- `container/list` maintains access order.
- `Get` promotes an entry to most-recently-used.
- `Set` updates an existing entry or evicts the least-recently-used entry at capacity.
- `Range` traverses values from least to most recently used.

## Test

```bash
go test ./...
```

## Origin

Implemented as part of the 2025 Yandex School of Data Analysis Go course, then extracted and documented as a standalone portfolio project. The original course repository is licensed under MIT; its license is included in this repository.

