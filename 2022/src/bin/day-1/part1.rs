pub fn heaviest_elf() { 
    let max = include_str!("input-1.txt")
        .split("\n\n")
        .map(|x| return x.split("\n").flat_map(|x| x.parse::<usize>()).sum::<usize>())
        .max();

    if max.is_some() {
        println!("Day 1, Excercise 1; Heaviest Elf: {}", max.unwrap())
    } else {
        panic!("Something went wrong.")
    }
}
