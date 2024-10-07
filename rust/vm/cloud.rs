use std::collections::HashMap;
use std::sync::{Mutex, Arc};
use std::error::Error;

#[derive(Debug, Clone)]
pub struct User {
    pub id: String,
    // Add other fields as necessary
}

pub trait UserService {
    fn create_user(&self, user: User) -> Result<(), Box<dyn Error>>;
    fn get_user_by_id(&self, id: &str) -> Result<Option<User>, Box<dyn Error>>;
    fn update_user(&self, user: User) -> Result<(), Box<dyn Error>>;
    fn delete_user(&self, id: &str) -> Result<(), Box<dyn Error>>;
}

pub struct InMemoryUserService {
    users: Arc<Mutex<HashMap<String, User>>>,
}

impl InMemoryUserService {
    pub fn new() -> Self {
        InMemoryUserService {
            users: Arc::new(Mutex::new(HashMap::new())),
        }
    }
}

impl UserService for InMemoryUserService {
    fn create_user(&self, user: User) -> Result<(), Box<dyn Error>> {
        let mut users = self.users.lock().unwrap();

        if users.contains_key(&user.id) {
            return Err("user already exists".into());
        }

        users.insert(user.id.clone(), user);
        Ok(())
    }

    fn get_user_by_id(&self, id: &str) -> Result<Option<User>, Box<dyn Error>> {
        let users = self.users.lock().unwrap();
        Ok(users.get(id).cloned())
    }

    fn update_user(&self, user: User) -> Result<(), Box<dyn Error>> {
        let mut users = self.users.lock().unwrap();

        if users.contains_key(&user.id) {
            users.insert(user.id.clone(), user);
            Ok(())
        } else {
            Err("user not found".into())
        }
    }

    fn delete_user(&self, id: &str) -> Result<(), Box<dyn Error>> {
        let mut users = self.users.lock().unwrap();

        if users.remove(id).is_some() {
            Ok(())
        } else {
            Err("user not found".into())
        }
    }
}
