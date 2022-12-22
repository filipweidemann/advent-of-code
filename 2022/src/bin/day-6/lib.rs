use std::collections::HashSet;

pub enum Marker {
    Signal,
    Message
}

impl Marker {
    fn sequence_length(&self) -> usize {
        match self {
            Marker::Signal => 4,
            Marker::Message => 14
        }
    }
}

pub fn find_signal_marker(input: &str, marker: &Marker) -> usize {
    let matched_index = input.as_bytes().windows(marker.sequence_length()).position(|win| HashSet::<_>::from_iter(win).len() == marker.sequence_length()).unwrap_or(0);

    if matched_index != 0 {
        return matched_index + marker.sequence_length()
    } else {
        return 0
    }
}

#[cfg(test)]
mod tests {
    use crate::lib::{find_signal_marker, Marker};

    #[test]
    fn test_signal() {
        const INPUT: &str = include_str!("input-test.txt");
        let marker_end = find_signal_marker(INPUT, &Marker::Signal);
        assert_eq!(marker_end, 10)
    }

    #[test]
    fn test_message() {
        const INPUT: &str = include_str!("input-test.txt");
        let marker_end = find_signal_marker(INPUT, &Marker::Message);
        assert_eq!(marker_end, 29)
    }
}
