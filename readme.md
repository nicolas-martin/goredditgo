##GoRedditGo##

This is a project created on a 6 hour train ride to teach myself GO.

GoRedditGo takes in as a parameter one or more subreddit, then it looks for the best post (based on number of posts and score). It then uses the Aylien's GO API to extract key words / hashtags and then publishes a tweet.


##Usage##
The default behaviour looks for a `key` file in the root folder in the following order.

```
consumerKey
consumerSecret
accessToken
applicationID
applicationKey
```

The first 3 keys are required for the twitter api and the last 2 is the for the Aylien text analysis API

Ultimately these could be moved in the config file as environement variables as well.
