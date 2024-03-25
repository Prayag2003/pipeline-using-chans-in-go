# Pipeline using Go { Prime Number Generator }

This Go program generates prime numbers using goroutines and channels. It employs a slow pipeline approach to generate prime numbers from random integers.

## Theory

- The program utilizes goroutines and channels to implement a pipeline for generating and filtering prime numbers.
- It generates random integers, checks whether they are prime, and then sends the prime numbers through the pipeline.
- The pipeline consists of three main stages: generating random integers, checking primality, and filtering prime numbers.

## Installation and Usage

1. **Install Go:**
   Ensure you have Go installed on your system. You can download it from [golang.org](https://golang.org/).

2. **Clone the Repository:**

   ```bash
   git clone https://github.com/your-username/prime-number-generator.git
   cd prime-number-generator
   ```

3. **Run the Program:**
   ```bash
    go run main.go
   ```

## Overview

The code demonstrates a concurrent pipeline to generate prime numbers from randomly generated integers. It involves the following key components:

- **`repeatFunc`**: Generates an infinite stream of random integers using a provided function.
- **`take`**: Controls the number of elements taken from a stream.
- **`isPrimeStream`**: Determines if a number is a prime number. Note: This function uses a slow method to check for primality.
- **`fanIn`**: Orchestrates multiple channels into a single channel, aggregating data from different goroutines.

## PipeLine Design and Description

### `repeatFunc`

- **Purpose**: Creates an infinite stream of random integers.
- **Parameters**:
  - `done <-chan K`: Signal channel to indicate completion.
  - `fn func() T`: Function that generates a random value of type `T`.
- **Returns**: Read-only channel (`<-chan T`) of generated random values.

### `take`

- **Purpose**: Controls the number of elements taken from a given stream.
- **Parameters**:
  - `done <-chan K`: Signal channel to indicate completion.
  - `stream <-chan T`: Input stream of values.
  - `num int`: Number of elements to extract.
- **Returns**: Read-only channel (`<-chan T`) with a limited number of elements.

### `isPrimeStream`

- **Purpose**: Checks if the generated integers are prime.
- **Parameters**:
  - `done <-chan int`: Signal channel to indicate completion.
  - `randStream <-chan int`: Input stream of random integers.
- **Returns**: Read-only channel (`<-chan int`) containing only prime numbers.

### `fanIn`

- **Purpose**: Aggregates multiple channels into a single channel.
- **Parameters**:
  - `done <-chan int`: Signal channel to indicate completion.
  - `channels ...<-chan T`: Multiple input channels to merge.
- **Returns**: Single channel (`<-chan T`) combining data from all input channels.

## Usage

1. **Generating Prime Numbers**:

   - To generate prime numbers, uncomment the relevant section in the `main` function.
   - Adjust the `take` function with the desired number of prime numbers.

2. **Optimizing Prime Calculation**:

   - The `isPrimeStream` function currently employs a slow prime checking method. For larger numbers, consider implementing more efficient prime checking algorithms.

3. **Concurrency Control**:
   - The program leverages Go's concurrency model to generate and process prime numbers concurrently. The number of goroutines used can be adjusted based on the available CPU cores (`runtime.NumCPU()`).

## Notes

- The `isPrimeStream` function's primality check method is not optimized for larger numbers and might impact performance.
- Consider implementing a more efficient prime checking algorithm for larger numbers.

---

Feel free to further elaborate on specific optimizations or details about prime number generation techniques based on your expertise or preferences!
