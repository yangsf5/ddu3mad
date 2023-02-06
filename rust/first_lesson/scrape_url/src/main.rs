use std::fs;

fn scrape_url() {
    let url = "https://www.rust-lang.org/";
    let output = "rust.md";

    println!("Fetching url: {}", url);
    let body = reqwest::blocking::get(url).unwrap().text().unwrap();

    println!("Converting html to markdown...");
    let md = html2md::parse_html(&body);

    fs::write(output, md.as_bytes()).unwrap();
    println!("Converted markdown has been saved in {}.", output);
}

fn apply(value: i32, f: fn(i32) -> i32) -> i32 {
    f(value)
}

fn square(value: i32) -> i32 {
    value * value
}

fn cube(value: i32) -> i32 {
    value * value * value
}

fn pi() -> f64 {
    3.1415926
}

fn not_pi() {
    3.1415926;
}

fn main() {
    // scrape_url();

    println!("apply square: {}", apply(3, square));
    println!("apply cube: {}", apply(3, cube));

    let is_pi = pi();
    let is_unit1 = not_pi();
    let is_unit2 = {
        pi();
    };

    println!(
        "is_pi: {:?}, is_unit1: {:?}, is_unit2: {:?}",
        is_pi, is_unit1, is_unit2
    );
}
