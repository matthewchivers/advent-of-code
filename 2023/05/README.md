# Day 5: If You Give A Seed A Fertilizer

You take the boat and find the gardener right where you were told he would be: managing a giant "garden" that looks more like a farm.

"A water source? Island Island is the water source!" You point out that Snow Island isn't receiving any water.

"Oh, we had to stop the water because we ran out of sand to filter it with! Can't make snow with dirty water. Don't worry, I'm sure we'll get more sand soon; we only turned off the water a few days... weeks... oh no." His face sinks into a look of horrified realisation.

"I've been so busy making sure everyone here has food that I completely forgot to check why we stopped getting more sand! There's a ferry leaving soon that is headed over in that direction - it's much faster than your boat. Could you please go check it out?"

You barely have time to agree to this request when he brings up another. "While you wait for the ferry, maybe you can help us with our food production problem. The latest Island Island Almanac just arrived and we're having trouble making sense of it."

The almanac (your puzzle input) lists all of the seeds that need to be planted. It also lists what type of soil to use with each kind of seed, what type of fertiliser to use with each kind of soil, what type of water to use with each kind of fertiliser, and so on. Every type of seed, soil, fertiliser, etc., is identified with a number, but numbers are reused by each category - that is, soil 123 and fertiliser 123 aren't necessarily related to each other.

### Example Inputs:

**Seeds:**
```
79 14 55 13
```

**Seed-to-Soil Map:**
```
50 98 2
52 50 48
```

**Soil-to-Fertiliser Map:**
```
0 15 37
37 52 2
39 0 15
```

**Fertiliser-to-Water Map:**
```
49 53 8
0 11 42
42 0 7
57 7 4
```

**Water-to-Light Map:**
```
88 18 7
18 25 70
```

**Light-to-Temperature Map:**
```
45 77 23
81 45 19
68 64 13
```

**Temperature-to-Humidity Map:**
```
0 69 1
1 0 69
```

**Humidity-to-Location Map:**
```
60 56 37
56 93 4
```

The almanac starts by listing which seeds need to be planted: seeds `79, 14, 55, and 13`.

The rest of the almanac contains a list of maps which describe how to convert numbers from a source category into numbers in a destination category. For instance, the section that starts with **seed-to-soil map** describes how to convert a seed number (source) to a soil number (destination). This lets the gardener know which soil to use with which seeds, which water to use with which fertiliser, and so on.

### Conversion Example:

Consider the **seed-to-soil map**:

```
50 98 2
52 50 48
```

- The first line (`50 98 2`) means that source range starts at 98 and contains two values (`98, 99`), and the destination range starts at `50` with two values (`50, 51`).

- The second line means that the source range starts at `50` with `48` values (`50, 51, ..., 96, 97`) and corresponds to a destination range starting at `52`.

Any source numbers that aren't mapped correspond to the same destination number. For example, seed number `10` corresponds to soil number `10`.

With this map, you can look up the soil number required for each initial seed number:

- **Seed number 79** corresponds to **soil number 81**.
- **Seed number 14** corresponds to **soil number 14**.
- **Seed number 55** corresponds to **soil number 57**.
- **Seed number 13** corresponds to **soil number 13**.

### Finding the Lowest Location Number

The gardener and his team want to start immediately, so they'd like to know the **closest location** that needs a seed. To determine this, convert each seed through the entire sequence of categories.

**Example Corresponding Types:**

- **Seed 79**: Soil 81 → Fertiliser 81 → Water 81 → Light 74 → Temperature 78 → Humidity 78 → Location 82.
- **Seed 14**: Soil 14 → Fertiliser 53 → Water 49 → Light 42 → Temperature 42 → Humidity 43 → Location 43.
- **Seed 55**: Soil 57 → Fertiliser 57 → Water 53 → Light 46 → Temperature 82 → Humidity 82 → Location 86.
- **Seed 13**: Soil 13 → Fertiliser 52 → Water 41 → Light 34 → Temperature 34 → Humidity 35 → Location 35.

Thus, the lowest location number in this example is **35**.

### Puzzle Question:

What is the lowest location number that corresponds to any of the initial seed numbers?

Your puzzle answer was **3374647**.

---

## Part Two

Everyone will starve if only a small number of seeds are planted. Upon re-reading, it looks like the **seeds** line actually describes **ranges of seed numbers**.

The values come in pairs. Within each pair, the first value is the start of the range, and the second value is the range length.

### Updated Example:

**Seeds:**
```
79 14 55 13
```
This describes **two ranges** of seed numbers:

- **Range 1**: Starts at `79`, contains `14` values (`79, 80, ..., 91, 92`).
- **Range 2**: Starts at `55`, contains `13` values (`55, 56, ..., 66, 67`).

Now, rather than considering four seed numbers, you need to consider **27 seed numbers**.

**Example Lowest Location:**

In this updated scenario, the **lowest location number** is `46`.

### Puzzle Question:

What is the lowest location number that corresponds to any of the initial seed numbers?

Your puzzle answer was **6082852**.
