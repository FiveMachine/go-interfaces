use std::sync::Arc;
use async_trait::async_trait;
use anyhow::Result;
use tokio::sync::RwLock;
use std::collections::HashMap;

/// Memory trait representing a module's memory.
pub trait Memory: Send + Sync {
    fn size(&self) -> u32;
    fn grow(&mut self, delta_pages: u32) -> Result<u32>;
}

/// Function trait representing a function exported from the module.
pub trait Function: Send + Sync {
    fn call(&self, params: &[u64]) -> Result<Vec<u64>>;
}

/// Global trait representing a global variable exported from the module.
pub trait Global: Send + Sync {
    fn get(&self) -> u64;
    fn set(&mut self, value: u64);
}

/// Module trait representing a post-instantiated module.
#[async_trait]
pub trait Module: Send + Sync {
    /// Returns the name the module was instantiated with.
    fn name(&self) -> String;

    /// Returns the memory defined in this module, if any.
    fn memory(&self) -> Option<Arc<dyn Memory>>;

    /// Returns a function exported from this module by name, if any.
    fn exported_function(&self, name: &str) -> Option<Arc<dyn Function>>;

    /// Returns a memory exported from this module by name, if any.
    fn exported_memory(&self, name: &str) -> Option<Arc<dyn Memory>>;

    /// Returns a global exported from this module by name, if any.
    fn exported_global(&self, name: &str) -> Option<Arc<dyn Global>>;

    /// Closes the module with a specified exit code.
    async fn close_with_exit_code(&self, exit_code: u32) -> Result<()>;

    /// Closes the module, releasing any resources.
    async fn close(&self, ctx: Option<tokio::sync::oneshot::Receiver<()>>) -> Result<()>;
}

/// Example implementation of the `Module` trait.
pub struct SimpleModule {
    name: String,
    memory: Option<Arc<dyn Memory>>,
    functions: HashMap<String, Arc<dyn Function>>,
    globals: HashMap<String, Arc<dyn Global>>,
}

impl SimpleModule {
    pub fn new(
        name: String,
        memory: Option<Arc<dyn Memory>>,
        functions: HashMap<String, Arc<dyn Function>>,
        globals: HashMap<String, Arc<dyn Global>>,
    ) -> Self {
        Self {
            name,
            memory,
            functions,
            globals,
        }
    }
}

#[async_trait]
impl Module for SimpleModule {
    fn name(&self) -> String {
        self.name.clone()
    }

    fn memory(&self) -> Option<Arc<dyn Memory>> {
        self.memory.clone()
    }

    fn exported_function(&self, name: &str) -> Option<Arc<dyn Function>> {
        self.functions.get(name).cloned()
    }

    fn exported_m
