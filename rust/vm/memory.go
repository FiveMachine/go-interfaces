use std::convert::TryInto;
use std::sync::Arc;
use anyhow::Result;

pub const MEMORY_PAGE_SIZE: u32 = 65536;
pub const MEMORY_LIMIT_PAGES: u32 = 65536;

/// Memory trait provides restricted access to a module's memory.
pub trait Memory: Send + Sync {
    /// Returns the size of the memory in pages.
    fn size(&self) -> u32;

    /// Grows the memory by `delta_pages` and returns the previous size in pages.
    fn grow(&mut self, delta_pages: u32) -> Result<u32>;

    /// Reads a single byte from the memory at the given offset.
    fn read_byte(&self, offset: u32) -> Option<u8>;

    /// Reads a `u16` in little-endian encoding from the memory at the given offset.
    fn read_u16_le(&self, offset: u32) -> Option<u16>;

    /// Reads a `u32` in little-endian encoding from the memory at the given offset.
    fn read_u32_le(&self, offset: u32) -> Option<u32>;

    /// Reads a `f32` in IEEE 754 little-endian encoding from the memory at the given offset.
    fn read_f32_le(&self, offset: u32) -> Option<f32>;

    /// Reads a `u64` in little-endian encoding from the memory at the given offset.
    fn read_u64_le(&self, offset: u32) -> Option<u64>;

    /// Reads a `f64` in IEEE 754 little-endian encoding from the memory at the given offset.
    fn read_f64_le(&self, offset: u32) -> Option<f64>;

    /// Reads `byte_count` bytes from the memory starting at the given offset.
    fn read(&self, offset: u32, byte_count: u32) -> Option<Vec<u8>>;

    /// Writes a single byte to the memory at the given offset.
    fn write_byte(&mut self, offset: u32, value: u8) -> bool;

    /// Writes a `u16` in little-endian encoding to the memory at the given offset.
    fn write_u16_le(&mut self, offset: u32, value: u16) -> bool;

    /// Writes a `u32` in little-endian encoding to the memory at the given offset.
    fn write_u32_le(&mut self, offset: u32, value: u32) -> bool;

    /// Writes a `f32` in IEEE 754 little-endian encoding to the memory at the given offset.
    fn write_f32_le(&mut self, offset: u32, value: f32) -> bool;

    /// Writes a `u64` in little-endian encoding to the memory at the given offset.
    fn write_u64_le(&mut self, offset: u32, value: u64) -> bool;

    /// Writes a `f64` in IEEE 754 little-endian encoding to the memory at the given offset.
    fn write_f64_le(&mut self, offset: u32, value: f64) -> bool;

    /// Writes a slice of bytes to the memory starting at the given offset.
    fn write(&mut self, offset: u32, values: &[u8]) -> bool;
}

/// A MemorySizer function type for resizing memory during compilation.
pub type MemorySizer = Arc<dyn Fn(u32, Option<u32>) -> (u32, u32, u32) + Send + Sync>;

/// Implementation of Memory trait using a Vec<u8> as the backing store.
pub struct SimpleMemory {
    buffer: Vec<u8>,
    size: u32, // Size in pages
}

impl SimpleMemory {
    pub fn new(initial_pages: u32, max_pages: Option<u32>) -> Self {
        let size_in_bytes = (initial_pages * MEMORY_PAGE_SIZE) as usize;
        SimpleMemory {
            buffer: vec![0; size_in_bytes],
            size: initial_pages,
        }
    }
}

impl Memory for SimpleMemory {
    fn size(&self) -> u32 {
        self.size
    }

    fn grow(&mut self, delta_pages: u32) -> Result<u32> {
        let previous_size = self.size;
        let new_size = self.size.checked_add(delta_pages).ok_or(anyhow::Error::msg("Memory overflow"))?;

        if new_size > MEMORY_LIMIT_PAGES {
            return Err(anyhow::Error::msg("Memory limit exceeded"));
        }

        let new_size_in_bytes = (new_size * MEMORY_PAGE_SIZE) as usize;
        self.buffer.resize(new_size_in_bytes, 0);
        self.size = new_size;

        Ok(previous_size)
    }

    fn read_byte(&self, offset: u32) -> Option<u8> {
        self.buffer.get(offset as usize).cloned()
    }

    fn read_u16_le(&self, offset: u32) -> Option<u16> {
        let slice = self.buffer.get(offset as usize..(offset as usize + 2))?;
        Some(u16::from_le_bytes(slice.try_into().ok()?))
    }

    fn read_u32_le(&self, offset: u32) -> Option<u32> {
        let slice = self.buffer.get(offset as usize..(offset as usize + 4))?;
        Some(u32::from_le_bytes(slice.try_into().ok()?))
    }

    fn read_f32_le(&self, offset: u32) -> Option<f32> {
        self.read_u32_le(offset).map(f32::from_bits)
    }

    fn read_u64_le(&self, offset: u32) -> Option<u64> {
        let slice = self.buffer.get(offset as usize..(offset as usize + 8))?;
        Some(u64::from_le_bytes(slice.try_into().ok()?))
    }

    fn read_f64_le(&self, offset: u32) -> Option<f64> {
        self.read_u64_le(offset).map(f64::from_bits)
    }

    fn read(&self, offset: u32, byte_count: u32) -> Option<Vec<u8>> {
        let end = offset.checked_add(byte_count)? as usize;
        Some(self.buffer.get(offset as usize..end)?.to_vec())
    }

    fn write_byte(&mut self, offset: u32, value: u8) -> bool {
        if let Some(byte) = self.buffer.get_mut(offset as usize) {
            *byte = value;
            true
        } else {
            false
        }
    }

    fn write_u16_le(&mut self, offset: u32, value: u16) -> bool {
        if let Some(slice) = self.buffer.get_mut(offset as usize..(offset as usize + 2)) {
            slice.copy_from_slice(&value.to_le_bytes());
            true
        } else {
            false
        }
    }

    fn write_u32_le(&mut self, offset: u32, value: u32) -> bool {
        if let Some(slice) = self.buffer.get_mut(offset as usize..(offset as usize + 4)) {
            slice.copy_from_slice(&value.to_le_bytes());
            true
        } else {
            false
        }
    }

    fn write_f32_le(&mut self, offset: u32, value: f32) -> bool {
        self.write_u32_le(offset, value.to_bits())
    }

    fn write_u64_le(&mut self, offset: u32, value: u64) -> bool {
        if let Some(slice) = self.buffer.get_mut(offset as usize..(offset as usize + 8)) {
            slice.copy_from_slice(&value.to_le_bytes());
            true
        } else {
            false
        }
    }

    fn write_f64_le(&mut self, offset: u32, value: f64) -> bool {
        self.write_u64_le(offset, value.to_bits())
    }

    fn write(&mut self, offset: u32, values: &[u8]) -> bool {
        let end = offset as usize + values.len();
        if end > self.buffer.len() {
            return false;
        }
        self.buffer[offset as usize..end].copy_from_slice(values);
        true
    }
}
