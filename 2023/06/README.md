# Day 6: Wait For It

> Note: This is one that required some thought and learning for me. My first attempt was brute-force, but that ended up with a Part 1 that took a while to compute, and a part 2 that never even finished for me - so I turned to _maths_!
> Scroll to the bottom of the instructions for my explanation:

## Instructions

The ferry quickly brings you across Island Island. After asking around, you discover that there is indeed normally a large pile of sand somewhere near here, but you don't see anything besides lots of water and the small island where the ferry has docked.

As you try to figure out what to do next, you notice a poster on a wall near the ferry dock. "Boat races! Open to the public! Grand prize is an all-expenses-paid trip to Desert Island!" That must be where the sand comes from! Best of all, the boat races are starting in just a few minutes.

You manage to sign up as a competitor in the boat races just in time. The organiser explains that it's not really a traditional race - instead, you will get a fixed amount of time during which your boat has to travel as far as it can, and you win if your boat goes the farthest.

As part of signing up, you get a sheet of paper (your puzzle input) that lists the time allowed for each race and also the best distance ever recorded in that race. To guarantee you win the grand prize, you need to make sure you go farther in each race than the current record holder.

The organiser brings you over to the area where the boat races are held. The boats are much smaller than you expected - they're actually toy boats, each with a big button on top. Holding down the button charges the boat, and releasing the button allows the boat to move. Boats move faster if their button was held longer, but time spent holding the button counts against the total race time. You can only hold the button at the start of the race, and boats don't move until the button is released.

For example:

| Time     | 7 | 15 | 30  |
| -------- | - | -- | --- |
| Distance | 9 | 40 | 200 |

This document describes three races:

- The first race lasts **7 milliseconds**. The record distance in this race is **9 millimetres**.
- The second race lasts **15 milliseconds**. The record distance in this race is **40 millimetres**.
- The third race lasts **30 milliseconds**. The record distance in this race is **200 millimetres**.

Your toy boat has a starting speed of zero millimetres per millisecond. For each whole millisecond you spend at the beginning of the race holding down the button, the boat's speed increases by one millimetre per millisecond.

So, because the first race lasts 7 milliseconds, you only have a few options:

- Don't hold the button at all (that is, hold it for 0 milliseconds) at the start of the race. The boat won't move; it will have travelled **0 millimetres** by the end of the race.
- Hold the button for 1 millisecond at the start of the race. Then, the boat will travel at a speed of **1 millimetre per millisecond** for **6 milliseconds**, reaching a total distance travelled of **6 millimetres**.
- Hold the button for 2 milliseconds, giving the boat a speed of **2 millimetres per millisecond**. It will then get **5 milliseconds** to move, reaching a total distance of **10 millimetres**.
- Hold the button for 3 milliseconds. After its remaining **4 milliseconds** of travel time, the boat will have gone **12 millimetres**.
- Hold the button for 4 milliseconds. After its remaining **3 milliseconds** of travel time, the boat will have gone **12 millimetres**.
- Hold the button for 5 milliseconds, causing the boat to travel a total of **10 millimetres**.
- Hold the button for 6 milliseconds, causing the boat to travel a total of **6 millimetres**.
- Hold the button for 7 milliseconds. That's the entire duration of the race. You never let go of the button. The boat can't move until you let go of the button. Please make sure you let go of the button so the boat gets to move. **0 millimetres**.

Since the current record for this race is **9 millimetres**, there are actually **4** different ways you could win: you could hold the button for **2, 3, 4, or 5** milliseconds at the start of the race.

In the second race, you could hold the button for at least **4 milliseconds** and at most **11 milliseconds** and beat the record, a total of **8** different ways to win.

In the third race, you could hold the button for at least **11 milliseconds** and no more than **19 milliseconds** and still beat the record, a total of **9** ways you could win.

To see how much margin of error you have, determine the number of ways you can beat the record in each race; in this example, if you multiply these values together, you get **288** (4 \* 8 \* 9).

Determine the number of ways you could beat the record in each race. What do you get if you multiply these numbers together?

Your puzzle answer was **6209190**.

## Part Two

As the race is about to start, you realise the piece of paper with race times and record distances you got earlier actually just has very bad kerning. There's really only one race - ignore the spaces between the numbers on each line.

So, the example from before:

| Time     | 7 | 15 | 30  |
| -------- | - | -- | --- |
| Distance | 9 | 40 | 200 |

...now instead means this:

| Time     | 71530  |
| -------- | ------ |
| Distance | 940200 |

Now, you have to figure out how many ways there are to win this single race. In this example, the race lasts for **71530 milliseconds** and the record distance you need to beat is **940200 millimetres**. You could hold the button anywhere from **14** to **71516 milliseconds** and beat the record, a total of **71503** ways!

How many ways can you beat the record in this one much longer race?

Your puzzle answer was **28545089**.

# Methodology

Though it turns out that (in adulthood) I do actually enjoy maths and somewhat excel in it, it's not something I have historically been acclaimed for. So this one took a bit of googling for me.

Task: Find the number of values/solutions that create a race distance longer than the current record. 

It turns out that, to find a range of values that satisfy an inequality, we can use quadratic equations!

The explanations below might be a bit verbose for those of you with degrees or A-Levels in mathematics - but I approached this having never properly studied quadratic equations before now, so the low-level explanation was useful for me.

## Understanding the problem:

We have a few key points about how the boat/race works:

1. Charging Time (`t`): At the start of the race, press a button th charge the boat for `(t)` milliseconds.
1. Speed: The boat's eventual speed increases the longer the button is held. Holding the button for `t` milliseconds sets the boat's speed to `t` millimetres per millisecond.
1. Total race time (`T`): The total time in which a boat may charge and/or move.
1. Travel time (`T - t`): After charging, the boat gets to move for the remaining time.
1. Distance travelled (`S`): The boat's total distance is calculated by multiplying the speed (`t mm/ms`) by the time left to move (`T - t` milliseconds):
``` math
S = t * (T - t)
```

The broad goal is to work out how many values of `t` would let the boat travel farther than the record distance of each race (`D` is the record distance / the one to beat).

In other words, I needed to work out how to make:
``` math
S > D
```
or 
``` math
 t * (T - t) > D
```
The distance travelled (`S` or `t * (T - t)`) needs to be greater than the record distance `D`.

So, we need to find all values of `t` that make `S` greater than `D`.

### How do we do this programatically?

#### Step 1 - Expand the statement 

Multiply everything out of brackets (one one side of the statement) by everything inside brackets (on the same side of the statement):

``` math
t * (T - t) > D
```

We do `t * T` giving `tT`, and `t * -t` to make `-t^2`.

``` math
t * (T - t) = tT - t^2
```

Now our inequality looks like:
``` math
tT - t^2 > D
```

#### Step 2 - Bring all terms to one side:

Subtract D from both sides to make:
``` math
tT - t^2 - D > 0
```

We do this just to make it a bit easier to understand and see how it behaves.

#### Step 3 - Rearrange the terms

It turns out that a quadratic equation is just one where `t` (or some other value) is multiplied by itself (`t ^ 2`).

A standard quadratic equation goes in the order of powers of `t`, so we rewrite ours as:

``` math
-t^2 + tT - D > 0
```

At this point, we then decide that it's easier to work with a positive `t^2` term, so multiply _everything_ by `-1`.

``` math
t^2 - tT + D < 0
```

Now we have a standard quadratic form, with a positive `t^2` and we need to find out where this expression is less than zero.

### Using the Quadratic Formula

To solve this, I used "The Quadratic Formula", which is a way to find he points (roots) where a quadratic espression equals zero. The formula is:

``` math
t = {-b ± \sqrt{b² - 4ac} \over 2a}
```

In our inequality we have:
- `a = 1` - coefficient of `t^2` (`1t ^ 2`)
- `b = -T` - coefficient of `t` (`-tT = -T * t`)
- `c = D` - constant term (number without variable - the distacne we're trying to beat)

Substituting these values into the formula:

1. **Calculate `b\):**

``` math
-b = -(-T) = T
```

2. **Calculate the Discriminant (Delta):**

The discriminant is the part under the square root:

``` math
\Delta 
```
can also be written as
``` math
b^2 - 4ac 
```
or
``` math
(-T)^2 - 4(1)(D) 
```
or 
``` math
T^2 - 4D
```

3. **Calculate the Roots:**

Using the quadratic formula:

``` math
t = \frac{T \pm \sqrt{T^2 - 4D}}{2}
```

This gives two roots:

``` math
t_1 = \frac{T - \sqrt{T^2 - 4D}}{2}
```

``` math
t_2 = \frac{T + \sqrt{T^2 - 4D}}{2}
```

### Finding the Range of Valid `t`

The roots `t_1`​ and `t_2`​ are the points where the quadratic expression equals zero, which divides the number line (above) into three sections. Because the quadratic curve opens upwards on a graph (like a 'U' shape), the expression is negative between the two "roots", meaning the graph dips _below_ the x-axis in this range. Therefore, if you are trying to determine where the expression is less than zero, the valid values of `t` are those that fall between `t_1`​ and `t_2`​.

``` math
t_1 < t < t_2
```

So, the valid values of `t` are all the integers between `t_1` and `t_2`.

### Practical Constraints

To get an answer that's actually useful to us, we need to apply some practical constraints:

1. `t` must be non-negative (it represents milliseconds to hold down the button, which cannot be negative)
1. `t` must be less than the total race time `T`.
1. We actually need to Count the Integer Values Between the Roots. The number of valid values is given by:

``` math
   \text{Number of ways} = \lfloor t_2 \rfloor - \lceil t_1 \rceil + 1
```

The floor of `t_2` minus the ceiling of `t_1` + 1 is the number of ways this can be solved/beaten.

### Example

Let’s apply this to the first race:

- Total race time: `T = 7`
- Record distance: `D = 9`

1. **Calculate the Discriminant:**

``` math
\Delta = T^2 - 4D = 7^2 - 4(9) = 49 - 36 = 13
```

2. **Find the Roots:**

``` math
   t_1 = \frac{7 - \sqrt{13}}{2} \approx \frac{7 - 3.6056}{2} \approx 1.697
```

``` math
   t_2 = \frac{7 + \sqrt{13}}{2} \approx \frac{7 + 3.6056}{2} \approx 5.303
```

3. **Determine the Valid Range:**

   Using the floor and ceiling functions:

``` math
   \text{Start} = \lceil t_1 \rceil = \lceil 1.697 \rceil = 2
```

``` math
   \text{End} = \lfloor t_2 \rfloor = \lfloor 5.303 \rfloor = 5
```

   So the valid values of `t` are:

``` math
t = 2, 3, 4, 5
```

4. **Count the Number of Ways:**

``` math
   \text{Number of ways} = \text{End} - \text{Start} + 1 = 5 - 2 + 1 = 4
```

---

### Final Thoughts

This process gave me the valid range of `t` values for each race, and I repeated these steps for every race. For Part One, I multiplied the number of valid `t` values for each race to get the final answer. For Part Two, since there was only one race, the result was simply the number of valid `t` values for that single race.

