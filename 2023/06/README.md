# 2023 Day 6: If You Give A Seed A Fertilizer

Using puzzle from the [website 2023/day/6](https://adventofcode.com/2023/day/6).


# Day 6: Wait For It

> Note: This is one that required some thought and learning for me. My first attempt was brute-force, but that ended up with a Part 1 that took a while to compute, and a part 2 that never even finished for me - so I turned to _maths_!
> Scroll to the bottom of the instructions for my explanation:

In the boat races, the goal is to beat the record distance by strategically holding the button to charge the toy boat's speed. The first task involves determining how many ways you can win each race by holding the button for different durations, then multiplying the number of winning strategies for all races (e.g., 288 in the example). The second task reinterprets the input as a single long race, requiring you to calculate the total number of ways to beat the record for that single race (e.g., 71,503 in the example).

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

