use std::ops::{RangeInclusive};

enum InclusionType {
    FullInclusion,
    PartialInclusion
}

fn check_range_inclusion(ranges: (RangeInclusive<i32>, RangeInclusive<i32>), inclusion: InclusionType) -> bool {
    let r1 = ranges.0;
    let r2 = ranges.1;
    
    match inclusion {
        InclusionType::FullInclusion => {
            return 
                r1.contains(&r2.start()) && r1.contains(&r2.end()) || 
                r2.contains(&r1.start()) && r2.contains(&r1.end());
        },

        InclusionType::PartialInclusion => {
            return 
                (r1.contains(&r2.start()) || r1.contains(&r2.end())) || 
                (r2.contains(&r1.start()) || r2.contains(&r1.end()));
        }
    }

}

fn check_partial_range_inclusion(ranges: (RangeInclusive<i32>, RangeInclusive<i32>)) {
    let r1 = ranges.0;
    let r2 = ranges.1;

    return 
}

fn ranges_from_tuple_pairs(boundaries: Vec<(i32, i32)>) -> (RangeInclusive<i32>, RangeInclusive<i32>) {
    return (
        (boundaries[0].0..=boundaries[0].1),
        (boundaries[1].0..=boundaries[1].1),
    )
}

fn get_ranges(input: &str) -> Vec<(RangeInclusive<i32>, RangeInclusive<i32>)> {
    let ranges: Vec<(RangeInclusive<i32>, RangeInclusive<i32>)> = input
        .lines()
        .map(|l| l.split(",")
             .filter_map(
                |c| 
                    c.split_once("-")
                    .and_then(
                        |(l, r)| Some(
                            (l.parse::<i32>().ok().unwrap(), 
                             r.parse::<i32>().ok().unwrap()
                            )
                        ) 
                    ) 
                )
             .collect::<Vec<(i32, i32)>>()
            )
        .map(ranges_from_tuple_pairs)
        .collect::<Vec<(RangeInclusive<i32>, RangeInclusive<i32>)>>();

    return ranges
}

pub fn count_full_overlaps(input: &str) -> usize {
    return get_ranges(input)
        .into_iter()
        .map(|ranges| check_range_inclusion(ranges, InclusionType::FullInclusion))
        .filter(|inclusion| *inclusion).count()
}

pub fn count_partial_overlaps(input: &str) -> usize {
    return get_ranges(input)
        .into_iter()
        .map(|ranges| check_range_inclusion(ranges, InclusionType::PartialInclusion))
        .filter(|inclusion| *inclusion).count()
}


#[cfg(test)]
mod tests {
    use crate::lib::{count_full_overlaps, count_partial_overlaps};

    use super::check_range_inclusion;

    
    #[test]
    fn test_check_range_inclusion() {
        assert!(check_range_inclusion(((1..=5), (2..=4)), super::InclusionType::FullInclusion))
    }

    #[test]
    fn test_count_full_overlaps() {
        const INPUT: &str = include_str!("input-test.txt");
        assert_eq!(count_full_overlaps(INPUT), 2);
    }

    #[test]
    fn test_count_partial_overlaps() {
        const INPUT: &str = include_str!("input-test.txt");
        assert_eq!(count_partial_overlaps(INPUT), 4)
    }
}
