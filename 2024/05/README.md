# Advent of Code 2024 - Day 5

In this readme I detail the puzzle instructions from [adventofcode.com/2024/day/5](https://adventofcode.com/2024/day/5) along with the answers to the puzzle, and any explanation on how I solved it. I typically try to comment the code in such a way that further explanation is not needed, but if things are particularly peculiar then I'll write a little something here.

## Instruction Summary - Day 5: Print Queue

The task here is to help an Elf in the North Pole printing department ensure that safety manual updates are printed in the correct order. Each page update must follow specific ordering rules, represented as pairs (e.g., `X|Y` means page `X` must be printed before page `Y` if both are included). Updates consist of lists of page numbers.

In **Part One**, the goal is to figure out which updates already meet the ordering requirements and identify the middle page numbers for those updates. Adding up these middle numbers provides the solution for this part.

In **Part Two**, the updates that aren’t in the correct order need to be rearranged to satisfy the rules. Once reordered, the middle page numbers are recalculated, and their sum provides the solution for this part.

## Approach

### Solving Day 5: Print Queue with Kahn's Algorithm

Part One was simple enough: loop through the updates and check to see if each one follows the specified rules. If it did, I grabbed the middle page number and added it to the total. This only needed a straightforward validation function to ensure the pages were in the correct order.

Part Two was more challenging since I had to reorder updates that didn’t follow the rules. Brute-forcing combinations was too slow, and naive sorting couldn’t handle the complex dependencies. A quick Google led me to Kahn’s Algorithm, a method for organising items into an order that respects their dependencies ("rules" in this case).

Here’s how I tackled it:

1. **Turn the Problem Into a Graph**
   Each page is a node, and the rules (e.g., `X|Y`) are edges between them. If `X|Y` exists, X has to come before Y. The goal is to sort the pages so all these rules are respected.

2. **Build the Graph**
   I pulled out the rules relevant to the current update and built:

   - An **adjacency list** to track which pages depend on which.
   - An **in-degree map** to count how many prerequisites each page has. Pages with no prerequisites (in-degree 0) go first.

3. **Set Up the Queue**
   Any page with in-degree 0 goes straight into the queue. These are the pages we can print right away without breaking any rules.

4. **Process the Queue**
   While the queue isn’t empty, I:

   - Remove a page from the queue (a page that’s ready to print).
   - Add it to the sorted list.
   - Update the in-degrees of its dependents. If any of them drop to 0, they get added to the queue.

5. **Check for Cycles**
   If I don’t end up with all the pages in the sorted list, it means there’s a cycle (like conflicting rules). Thankfully, that didn’t happen in my input.

### Final Thoughts

Using Kahn’s Algorithm, I was able to reorder the messy updates in a way that respected all the rules. Once sorted, calculating their middle page numbers and adding them up was straightforward. The algorithm proved to be a reliable tool for figuring out dependencies and solving the problem efficiently.
