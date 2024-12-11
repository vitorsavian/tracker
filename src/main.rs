mod cmd;
mod controller;
mod domain;
mod infra;

use crate::cmd::cli;

fn main() {
    cli::run();
}
