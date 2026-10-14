use std::collections::{HashMap,HashSet};

struct Twitter {
    followers: HashMap<i32,HashSet<i32>>,
    feed: HashMap<i32,Vec<(usize,i32)>>,
    tick:usize
}

impl Twitter {

    fn new() -> Self {
        Self { followers: HashMap::new(), feed: HashMap::new(), tick: 0 }
    }
    
    fn post_tweet(&mut self, user_id: i32, tweet_id: i32) {
        self.tick += 1;
        self.feed.entry(user_id).or_default().push((self.tick, tweet_id));       
    }
    
       fn get_news_feed(&self, user_id: i32) -> Vec<i32> {
        let mut ret = Vec::new();
        if let Some(tweets) = self.feed.get(&user_id) {
            ret.extend(tweets.iter().copied());
        }

        if let Some(followees) = self.followers.get(&user_id) {
            for followee in followees {
                if let Some(tweets) = self.feed.get(followee) {
                    ret.extend(tweets.iter().copied());
                }
            }
        }
        ret.sort_unstable_by(|a, b| b.0.cmp(&a.0));
        ret.into_iter().take(10).map(|(_, tweet_id)| tweet_id).collect()
    }
    
    fn follow(&mut self, follower_id: i32, followee_id: i32) {
        self.followers.entry(follower_id).or_default().insert(followee_id);
    }

    fn unfollow(&mut self, follower_id: i32, followee_id: i32) {
        if let Some(followees) = self.followers.get_mut(&follower_id) {
            followees.remove(&followee_id);
        }
    }
}