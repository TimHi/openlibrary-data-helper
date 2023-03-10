# Open Library Data Parser

Program to parse the provided [data dumps](https://openlibrary.org/developers/dumps) into a sqlite database.  
Since I created this for personal use I focused on some parts that are more important to me.

## Usage

Clone the repository, ensure go is installed and run `go run main.go`.  
Available options are:

```
  -rating string
        path to the rating data dump .txt
  -reading string
        path to the reading data dump .txt
  -transform string
        operation to apply on the data, available options: [top100]
```

## Whats done

| Type          | Status |
| ------------- | ------ |
| Ratings       | ✅     |
| Readings      | ✅     |
| Editions      | ❌     |
| Works         | ❌     |
| Authors       | ❌     |
| All Types     | ❌     |
| Complete Dump | ❌     |

## Transform

Here starts my personal use case, I wanted to extract the top 100 books, I did so by filtering for books that are read and have a few reviews.
Afterwards the results are enriched with data from the `/work` endpoint and stored as json.

## Performance

Could be better. Using Gorm it inserts faster.

## Issues

I'm not checking for existing entries. Duplicates may occure, I recommend running each dump only once. Checking for existing entries would make the performance even worse.

## Future Work

The database could be modeled smarter.
