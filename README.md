# **gator**

**gator** is a Blog Aggregator that allows users to follow their favorite RSS feeds!

Users can add and follow as many feeds as they like. **gator** then collects and stores posts from these feed in a PostgreSQL database. Users are then able to view summaries in the command line and access the full post via the provided URL. 

## Install

**gator** requires Go and Postgres

Make sure you have the latest [Go toolchain](https://golang.org/dl/) installed as well as a local Postgres database.

**gator** can be installed as a CLI command using 

```bash 
go install ...
```

## Config

Create a `.gatorconfig.json` file in your home directory with the following structure:

```json
{
  "db_url": "postgres://username:@localhost:5432/database?sslmode=disable"
}
```

Replace the values with your database connection string.

## Using **gator**
### Users
- gator register [username]: register a new user
- gator login [username]: login a registered user
- gator users: view currently registered users

### Feeds
- gator addfeed [title] [url]: add a new feed to the database (active user automatically follows added feed)
- gator feeds: display feeds in the database
- gator follow [url]: follow an existing feed in the database
- gator unfollow [url]: unfollow a feed
- gator following: view currently followed feeds
- gator agg [time]: collect posts from the internet and store them in the database
- gator browse [x *optional*]: preview x most recent feeds in command line

### Database
- reset: clears all users, feeds, and posts from the database


