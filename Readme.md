# Open Library Data Parser

Program to parse the provided [data dumps](https://openlibrary.org/developers/dumps) into a sqlite database.  
Since I created this for personal use I focused on some parts that are more important to me.

## Whats done

| Type  | Status  |  
|---|---|
| Ratings  |  ✅ |
| Readings  |  ✅ |
| Editions  |  ❌ |
| Works  | ❌  |
| Authors  | ❌  |
| All Types  | ❌  |
| Complete Dump  |  ❌ |

## Performance

Terrible. Using SQLC apparently has no bulk insert for SQLite. At least I wasnt able to find some. Feel free to open a PR to improve this.

## Issues

I'm not checking for existing entries. Duplicates may occure, I recommend running each dump only once. Checking for existing entries would make the performance even worse.
Also there is no brainwork done to handle larger files. The work or author dump should be a problem.