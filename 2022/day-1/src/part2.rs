use std::cmp::Reverse;

pub fn heaviest_elves(limit: usize) { 
    let mut maximums = include_str!("input-1.txt")
        .split("\n\n")
        .map(|x| return x.lines().flat_map(str::parse::<usize>).sum::<usize>()).collect::<Vec<usize>>();

    maximums.sort_by_key(|i| Reverse(*i));
    maximums.truncate(limit);
    
    println!("Day 1, Excercise 2; N-Heaviest Elves: {}", maximums.iter().sum::<usize>());
}
