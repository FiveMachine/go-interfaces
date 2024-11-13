pub mod vm {
    // Config sets configuration of the VM service
    pub struct Config {
        pub memory_limit_pages: u32, // should default to MEMORY_LIMIT_PAGES
        pub output: OutputType,
    }

    #[derive(Debug, Clone, Copy, PartialEq, Eq)]
    pub enum OutputType {
        Pipe,
        Buffer,
    }

    impl Default for Config {
        fn default() -> Self {
            Self {
                memory_limit_pages: MEMORY_LIMIT_PAGES, // assumes MEMORY_LIMIT_PAGES is defined
                output: OutputType::Pipe,
            }
        }
    }

    // Define the MEMORY_LIMIT_PAGES constant if needed
    pub const MEMORY_LIMIT_PAGES: u32 = 1024; // Set to your default limit
}
