# superperm
> An attempt to find an optimal heuristic solution to the superpermutation problem.

![license](https://img.shields.io/badge/License-MIT-lightgrey.svg?style=for-the-badge)
![version](https://img.shields.io/badge/Version-2.0.0-lightgrey.svg?style=for-the-badge)

## About

This project takes a heuristic approach when attempting to solve the superpermutation problem. The superpermutation problem is an open mathematics problem. At the moment, when the alphabet cardinality is 6, this algorithm does not find the shortest _known_ superpermutation. This project is still a work in progress and further attempts to optimize the algorithm will be made.

## Findings

| **\|alphabet\|** | **\|shortest(alphabet)\|** | **\|rotate(alphabet)\|** | **runtime(rotate(alphabet))** |
|:----------------:|:------------------------:|:---------------:|:---------------:|
| 1 | 1      | 1      | 750ns        |
| 2 | 3      | 3      | 2.166µs      |
| 3 | 9      | 9      | 3.666µs      |
| 4 | 33     | 33     | 8.583µs      |
| 5 | 153    | 153    | 47.25µs      |
| 6 | 872    | 873    | 254.709µs    |
| 7 | 5907   | 5913   | 2.402625ms   |
| 8 | 46205  | 46233  | 15.063542ms  |
| 9 | 408966 | 409113 | 138.998333ms |

## Development

Run `make help` for all available commands. In general, you can run `make build-all` to build the binary for all platforms.

## Additional Resources

- https://oeis.org/A180632
- https://en.wikipedia.org/wiki/Superpermutation
