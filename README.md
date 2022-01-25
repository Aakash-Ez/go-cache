# Go Cache
This repo is a TTLCache implementation in GoLang. The following assumptions have been made regarding the Cache Implementation:

- There is no limit on Cache Size
- When a new entry's key already exist in the Cache, the value is replaced. 

# Progress

| S. No | Progress                                     | Date        |
|-------|----------------------------------------------|-------------|
| 1     | Implemented using forever loop in go routine | 24 Jan 2022 |
| 2     | Implemented using time.Timer                 | 25 Jan 2022 |
