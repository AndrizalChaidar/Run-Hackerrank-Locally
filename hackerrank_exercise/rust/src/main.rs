use std::{
    cell::RefCell,
    collections::{HashMap, HashSet},
    io::{self, BufRead},
};
#[derive(Debug, Default)]
struct IdxVal {
    idx: Vec<u64>,
    val: Vec<u64>,
}

impl IdxVal {
    fn add_data(&mut self, idx: u64, health: u64) {
        self.idx.push(idx);
        self.val.push(health);
    }
}

fn main() {
    let stdin = io::stdin();
    let mut stdin_iterator = stdin.lock().lines();

    let n = stdin_iterator
        .next()
        .unwrap()
        .unwrap()
        .trim()
        .parse::<u64>()
        .unwrap();
    let mut data_map = HashMap::<&str, IdxVal>::default();
    let mut char_path = HashSet::<String>::new();

    let genes_string = stdin_iterator.next().unwrap().unwrap();
    let mut genes = genes_string.trim_end().split(' ');

    let health_string = stdin_iterator.next().unwrap().unwrap();
    let mut healths = health_string
        .trim_end()
        .split(' ')
        .map(|x| x.parse::<u64>().unwrap());

    for i in 0..n {
        let gene = genes.next().unwrap();
        let health = healths.next().unwrap();
        let gene_len = gene.len();
        let mut strings = String::new();
        let mut chars = gene.chars();
        for gene_idx in 0..gene_len {
            let ch = chars.next().unwrap();
            strings.insert(strings.len(), ch);
            char_path.insert(strings.clone());
            if gene_idx == gene_len - 1 {
                if let Some(data) = data_map.get_mut(&gene) {
                    data.add_data(i, health)
                } else {
                    data_map.insert(
                        gene,
                        IdxVal {
                            idx: vec![i],
                            val: vec![health],
                        },
                    );
                }
            }
        }
    }

    let s = stdin_iterator
        .next()
        .unwrap()
        .unwrap()
        .trim()
        .parse::<u32>()
        .unwrap();
    let mut max = u64::MIN;
    let mut min = u64::MAX;
    for _ in 0..s {
        let first_multiple_input: Vec<String> = stdin_iterator
            .next()
            .unwrap()
            .unwrap()
            .split(' ')
            .map(|s| s.to_string())
            .collect();

        let first = first_multiple_input[0].trim().parse::<u64>().unwrap();

        let last = first_multiple_input[1].trim().parse::<u64>().unwrap();

        let dna: Vec<String> = first_multiple_input[2]
            .chars()
            .enumerate()
            .map(|(_, x)| x.to_string())
            .collect();
        let res = RefCell::new(0 as u64);
        let add_res = |idxs: &Vec<u64>, values: &Vec<u64>| {
            for i in 0..idxs.len() {
                let idx = idxs[i];
                // println!("{:?}", idx);
                if idx >= first && idx <= last {
                    *res.borrow_mut() += values[i];
                }
            }
        };
        let dna_len = dna.len();
        for i in 0..dna_len {
            let mut chars = dna[i].clone();
            if let None = char_path.get(&chars) {
                continue;
            }
            if let Some(data) = data_map.get(chars.as_str()) {
                add_res(&data.idx, &data.val)
            }
            for j in (i + 1)..dna_len {
                chars.insert_str(chars.len(), dna[j].as_str());
                if let None = char_path.get(&chars) {
                    break;
                }
                if let Some(data) = data_map.get(chars.as_str()) {
                    add_res(&data.idx, &data.val)
                }
            }
        }
        let res_f = *res.borrow();

        if res_f > max {
            max = res_f
        }
        if res_f < min {
            min = res_f
        }
    }
    println!("{:?} {:?}", min, max);
}
