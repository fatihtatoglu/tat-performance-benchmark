# Huffman Coding

The Huffman code calculation contains steps to reach the result file.

## Normal Steps

- **Creating Frequency Table:** Calculate the existence of each char in the documents.
- **Creating Huffman Tree:** Creating a tree by using a frequency table. This table represents the shortcodes for each char. The most used chars will have shorter codes.
- **Creating Huffman Code:** Due to the table create the Huffman code for each char.
- **Encode the Text:** Rewrite the input document by replacing chars into the Huffman codes and providing compression.

## Implementing Steps

- Create the frequency table
  - Sort in descending order.
  - Persist the frequency table as a file (***filename***.freq).
- Create Huffman Tree
- Create Huffman Codes
  - Persist the codes as a file (***filename***.code).

## Input Files

- **Book1:** Moby Dick; Or, The Whale
- **Book2:** Oliver Twist
- **Book3:** A Tale of Two Cities
- **Book4:** Great Expectations

## Reference

- [Huffman coding](https://en.wikipedia.org/wiki/Huffman_coding)
- [Project Gutenberg](https://www.gutenberg.org/)
