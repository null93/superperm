# superperm
> An attempt to find an optimal heuristic solution to the superpermutation problem.

![license](https://img.shields.io/badge/License-MIT-lightgrey.svg?style=for-the-badge)
![version](https://img.shields.io/badge/Version-2.0.0-lightgrey.svg?style=for-the-badge)

<img width="882" height="617" alt="Screenshot" src="https://github.com/user-attachments/assets/cd20782b-e754-418c-94e4-e4683e4421e2" />

## About

This project takes a heuristic approach when attempting to solve the superpermutation problem. The superpermutation problem is an open mathematics problem. The `standard` strategy is the classic recursive construction with length 1! + 2! + ... + n!. Once the alphabet cardinality reaches 6, it does not find the shortest _known_ superpermutation. The `egan` strategy is Greg Egan's construction with length n! + (n-1)! + (n-2)! + (n-3)! + n - 3, which is shorter than `standard` once the alphabet cardinality reaches 7. This project is still a work in progress and further attempts to optimize the algorithm will be made.

## Findings

| **\|alphabet\|** | **\|shortest(alphabet)\|** | **\|standard(alphabet)\|** | **\|egan(alphabet)\|** |
|:----------------:|:------------------------:|:-----------------------:|:-------------------:|
| 1 | 1      | 1      | 1      |
| 2 | 3      | 3      | 3      |
| 3 | 9      | 9      | 9      |
| 4 | 33     | 33     | 34     |
| 5 | 153    | 153    | 154    |
| 6 | 872    | 873    | 873    |
| 7 | 5907   | 5913   | 5908   |
| 8 | 46181  | 46233  | 46205  |
| 9 | 408731 | 409113 | 408966 |

## Development

Run `make help` for all available commands. In general, you can run `make build-all` to build the binary for all platforms.

## Additional Resources

- https://oeis.org/A180632
- https://en.wikipedia.org/wiki/Superpermutation
- https://github.com/jaypantone/superperm-upper-43-80
- https://github.com/rumstd/superperm-upper-43-80_fork
