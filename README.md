# **gator**

**gator** is a Blog Aggregator that allows users to follow their favorite RSS feeds!

Users can add and follow as many feeds as they like. **gator** then collects and stores posts from these feed in a PostgreSQL database. Users are then able to view summaries in the command line and access the full post via the provided URL. 

**gator** requires Postgres and Go


**gator** can be installed as a CLI command using `go install`

The available commands for GATOR are listed below:
### Users
- register [username]: register a new user
- login [username]: login a registered user
- users: view currently registered users

### Feeds
- addfeed [title] [url]: add a new feed to the database (active user automatically follows added feed)
- feeds: display feeds in the database
- follow [url]: follow an existing feed in the database
- unfollow [url]: unfollow a feed
- following: view currently followed feeds
- agg [time]: collect posts from the internet and store them in the database
- browse [x *optional*]: preview x most recent feeds in command line

### Database
- reset: clears all users, feeds, and posts from the database


