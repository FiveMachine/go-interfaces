use std::sync::Arc;
use async_trait::async_trait;

/// A trait representing the context interface in Rust
#[async_trait]
pub trait Context: Send + Sync {
    /// Returns the Rust async runtime context (`tokio::context` equivalent).
    fn context(&self) -> Arc<tokio::sync::RwLock<tokio::task::JoinHandle<()>>>;

    /// Returns the project ID.
    fn project(&self) -> String;

    /// Returns the application name, or an empty string if none exists.
    fn application(&self) -> String;

    /// Returns the resource ID being used.
    fn resource() -> ID(Clone);

    /// Implement  branch function aspects parsing`]!("{name}=Options...")
    }

    
    
    

