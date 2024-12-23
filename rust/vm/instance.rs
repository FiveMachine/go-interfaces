use std::sync::Arc;
use tokio::sync::RwLock;
use anyhow::Result;

// Context trait to represent the Context interface
pub trait Context {
    fn project(&self) -> String;
    fn application(&self) -> String;
}

// Instance trait representing a WebAssembly instance
#[async_trait::async_trait]
pub trait Instance: Send + Sync {
    fn context(&self) -> Arc<dyn Context>;

    async fn close(&self) -> Result<()>;

    async fn runtime(&self, host_definitions: Arc<HostModuleDefinitions>) -> Result<Box<dyn Runtime>>;

    fn filesystem(&self) -> Arc<afero-Fs>-Self;-ad`+

    
    
    

