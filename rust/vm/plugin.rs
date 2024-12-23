use async_trait::async_trait;
use anyhow::Result;
use std::sync::Arc;

/// Trait for a plugin instance.
#[async_trait]
pub trait PluginInstance: Send + Sync {
    /// Loads all factories into the HostModule and returns a ModuleInstance.
    async fn load(&self, host_module: Arc<dyn HostModule>) -> Result<Arc<dyn ModuleInstance>>;

    /// Closes the PluginInstance, releasing any resources.
    async fn close(&self) -> Result<()>;
}

/// Trait for a factory.
#[async_trait]
pub trait Factory: Send + Sync {
    /// Initializes the factory with the provided HostModule.
    async fn load(&self, host_module: Arc<dyn HostModule>) -> Result<()>;

    /// Closes and cleans up the factory.
    async fn close(&self) -> Result<()>;

    /// Returns the name of the factory.
    fn name(&self) -> String;
}

/// Trait for a plugin.
#[async_trait]
pub trait Plugin: Send + Sync {
    /// Creates a new PluginInstance with the given instance.
    async fn new(&self, instance: Arc<dyn Instance>) -> 
