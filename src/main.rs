mod controller;

use clap::Parser;

#[derive(Parser, Debug)]
#[command(version, about, long_about = None)]
struct Args {
    // Book name
    #[arg(short, long)]
    name: String,

    // Number of pages
    #[arg(short, long, default_value_t = 1)]
    page: u8,

    // String database
    #[arg(long)]
    database: String,
}

fn main() {
    let args = Args::parse();
}
