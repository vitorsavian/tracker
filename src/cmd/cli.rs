use clap::{Parser, Subcommand};

use crate::controller::{novel, server::App};

#[derive(Parser)]
#[command(version, about, long_about = None)]
struct Cli {
    #[command(subcommand)]
    command: Commands,
}

#[derive(Subcommand)]
enum Commands {
    #[command(about = "Add a novel/book", long_about = None)]
    Add {
        #[arg(short, long)]
        name: String,

        #[arg(long)]
        finished: bool,

        #[arg(short, long)]
        chapter: u32,

        #[arg(short, long)]
        database: String,
    },

    Remove {
        #[arg(short, long)]
        name: String,

        #[arg(long)]
        finished: bool,

        #[arg(short, long)]
        chapter: u32,

        #[arg(short, long)]
        database: String,
    },

    Update {
        #[arg(short, long)]
        name: String,

        #[arg(long)]
        finished: bool,

        #[arg(short, long)]
        chapter: u32,

        #[arg(short, long)]
        database: String,
    },

    Get {
        #[arg(long)]
        id: String,

        #[arg(short, long)]
        all: bool,

        #[arg(short, long)]
        database: String,
    },
}

pub fn run() {
    let cli = Cli::parse();

    match cli.command {
        Commands::Add {
            name,
            finished,
            chapter,
            database,
        } => {
            let control = novel::Controller {
                conn: novel::Controller::database_connection(database),
            };

            control.conn.begin_test_transaction()
            println!("{:?}", finished);
            println!("{:?}", chapter);
        }

        Commands::Get { id, all, database } => {
            println!("{:?}", id);
            println!("{:?}", all);
            println!("{:?}", database);
        }

        Commands::Remove {
            name,
            finished,
            chapter,
            database,
        } => {
            println!("{:?}", name);
            println!("{:?}", finished);
            println!("{:?}", chapter);
            println!("{:?}", database);
        }

        Commands::Update {
            name,
            finished,
            chapter,
            database,
        } => {
            println!("{:?}", name);
            println!("{:?}", finished);
            println!("{:?}", chapter);
            println!("{:?}", database);
        }
    }
}
