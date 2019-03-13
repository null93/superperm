# superpermutation
> An attempt to find an optimal heuristic solution to the superpermutation problem.

![version](https://img.shields.io/badge/License-MIT-lightgrey.svg?style=for-the-badge)
![license](https://img.shields.io/badge/Version-1.0.0-lightgrey.svg?style=for-the-badge)

## About

This project takes a heuristic approach when attempting to solve the superpermutation problem. The superpermutation problem is an open mathematics problem. At the moment, when the alphabet cardinality is 6, this algorithm does not find the shortest _known_ superpermutation. This project is still a work in progress and further attempts to optimize the algorithm will be made.

## Findings

| **\|alphabet\|** | **\|superpermutation\|** | **timing (ms)** |
|:----------------:|:------------------------:|:---------------:|
|         1        |             1            |        0.453785 |
|         2        |             3            |        0.225103 |
|         3        |             9            |        0.048298 |
|         4        |            33            |        0.121518 |
|         5        |            153           |        1.065073 |
|         6        |            873           |       14.740946 |
|         7        |           5913           |      436.643946 |
|         8        |           46233          |     30867.66657 |
|         9        |          409113          |  1922906.717888 |

## Tools & Commands

|   **command**   |                    **description**                   |
|:---------------:|:-----------------------------------------------------|
| `yarn install`  | Install project dependencies.                        |
| `yarn generate` | Calculate superpermutations and save them to `dist`. |
| `yarn test`     | Test validity of generated superpermutations.        |
| `yarn clean`    | Delete generated superpermutations.                  |

## Additional Resources

- https://oeis.org/A180632
- https://en.wikipedia.org/wiki/Superpermutation
