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

I'm not checking for existing tables when running the program. Delete the sqlite.db after runs.
