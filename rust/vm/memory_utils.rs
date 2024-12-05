pub mod vm {
    /// Encodes the input as an Externref (uintptr equivalent in Rust).
    pub fn encode_externref(input: usize) -> u64 {
        input as u64
    }

    /// Decodes the input as an Externref (uintptr equivalent in Rust).
    pub fn decode_externref(input: u64) -> usize {
        input as usize
    }

    /// Encodes the input as an i32.
    pub fn encode_i32(input: i32) -> u64 {
        input as u64
    }

    /// Encodes the input as an i64.
    pub fn encode_i64(input: i64) -> u64 {
        input as u64
    }

    /// Encodes the input as an f32.
    pub fn encode_f32(input: f32) -> u64 {
        input.to_bits() as u64
    }

    /// Decodes the input as an f32.
    pub fn decode_f32(input: u64) -> f32 {
        f32::from_bits(input as u32)
    }

    /// Encodes the input as an f64.
    pub fn encode_f64(input: f64) -> u64 {
        input.to_bits()
    }

    /// Decodes the input as an f64.
    pub fn decode_f64(input: u64) -> f64 {
        f64::from_bits(input)
    }
}
