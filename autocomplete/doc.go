/*
Package cliautocomplete provide tools to make it so your program can support autocomplete (of flags, commands, subcommands, etc)
in various command line shells, such as: bash. zsh, fish, etc.

Autocomplete

Autocomplete is supported in some command line shells.

Perhaps the most well known of these shells, that support autocomplete, is the bash shell.

On the bash prompt, pressing the [TAB] key at certain times will cause the bash shell to either
№1 autocomplete the program, program (sub)command, program flag, file, etc, or
№2 output a list of  candidate programs, program (sub)commands, program flags, files, etc.

Do note that for №2, it is possible that the bash shell will partially autocomplete (when it can), and the output the candidates.
(More about that later though.)

For example, typing the following at the bash shell:
	b[TAB][TAB]

Would output something such as:
	b2sum         bash          bg            bmptoppm      brushtopbm    bunzip2       bzdiff        bzgrep        bzmore
	base32        bashbug       bind          bootctl       bsd-from      busctl        bzegrep       bzip2         
	base64        bdftopcf      bioradtopgm   break         bsd-write     bzcat         bzexe         bzip2recover  
	basename      bdftruncate   bmptopnm      broadwayd     builtin       bzcmp         bzfgrep       bzless        

These are the list of all the candidate programs. Note that all these candidate programs start with a “b”.

Note that why all that was outputted, instead of autocompletion happening, is because it is ambiguous which program was wanted.
Since there are many different programs that start with a “b”.
And the bash shell doesn't know which one you wanted.

OK, so what happens, for example, if we now typed the following at the bash shell:
	ba[TAB][TAB]

Well, this will first cause the bash shell to partially autocomplete it to:
	bas

(I.e., adding an “s” after the “ba”.)

And then output the candidates:
	base32    base64    basename  bash      bashbug

First note that all 5 of those candidate programs (“base32”, “base64”, “basename”, “bash”, “bashbug”) were in our previous output of candidate programs.

So, all that has happened is the bash shell is narrowing down the candidates, as you provide it with more information.

So, what would happen if instead you typed:
	bz[TAB][TAB]

Well, that would output:
	bzcat         bzdiff        bzexe         bzgrep        bzip2recover  bzmore        
	bzcmp         bzegrep       bzfgrep       bzip2         bzless        

(I.e., all the programs that begin with “bz”.)

And if you then typed:
	bzd[TAB][TAB]

It would autocomplete to:
	bzdiff

(And note that it did not output any candidates in this case, since it was able to do a full autocomplete.)

This type of autocomplete of programs is automagic.
Command line shells, such as the bash shell, have built-in support for it.

I.e., your program does not need to do anything to support that syle of autocomplete.

However, the types of autocomplete (of flags, of commands, of subcommands, etc) is something your program needs to do something extra (such as integrate the tools provided by this package) to enable.
But more on that a little later.

Autocomplete Flags

Let's now look at another type of autocomplte.
The autocompletion of flags.

Let's look at what happen if we were to type the following at the bash shell:
	grep --[TAB][TAB]

The bash shell would respond by outputting something such as:
	--after-context=         --count                  --files-with-matches     --line-buffered          --perl-regexp
	--basic-regexp           --dereference-recursive  --files-without-match    --line-number            --regexp=
	--before-context=        --devices=               --fixed-strings          --line-regexp            --silent
	--binary                 --directories=           --help                   --max-count=             --unix-byte-offsets
	--binary-files=          --exclude=               --ignore-case            --no-filename            --version
	--byte-offset            --exclude-dir=           --include=               --no-messages            --with-filename
	--color                  --exclude-from=          --initial-tab            --null                   --word-regexp
	--colour                 --extended-regexp        --invert-match           --null-data              
	--context=               --file=                  --label=                 --only-matching          

Note what happeded here.

Bash listed candidate flags for the “grep” program.

That is a powerful feature for the end user!
And creates a better user experience (UX) for the command line interface!

Now what if we were to type the following at the bash shell:
	grep --c[TAB][TAB]

The bash shell would first partially autocomplete this to:
	grep --co

And then output the candidates:
	--color     --colour    --context=  --count

If we then typed the following at the bash shell:
	grep --cou

The bash shell would then autocomplete it to:
	grep --count

(And not output any candidates, since it was able to do a full autocomplete.)

Autocomplete Commands, Subcommands, Etc

Let's now look at another type of autocomplte.
The autocompletion of commands, subcommands, etc.

Let's look at what happens if we were to type the following at the bash shell:
	tar [TAB][TAB]

The bash shell would respond by outputting the following candidate commands:
	A  c  d  r  t  u  x

The idea of a program supporting commands, subcommands, etc is an old one.

And although it is definitely not the first program to support commands, subcommands, etc,
maybe the most popular program today to support commands, subcommands, etc is: git.

So, let's look at what happens if we were to type the following at the bash shell:
	git [TAB][TAB]

The bash shell would then output the candidate commands:
	add                  cherry-pick          fsck                 mergetool            replace              submodule 
	am                   clean                gc                   mv                   request-pull         subtree 
	annotate             clone                get-tar-commit-id    name-rev             reset                tag 
	apply                commit               grep                 notes                revert               verify-commit 
	archive              config               help                 pull                 rm                   whatchanged 
	bisect               describe             imap-send            push                 shortlog             worktree 
	blame                diff                 init                 rebase               show                 
	branch               difftool             instaweb             reflog               show-branch          
	bundle               fetch                interpret-trailers   relink               stage                
	checkout             filter-branch        log                  remote               stash                
	cherry               format-patch         merge                repack               status               

And if we then typed the following at the bash shell:
	git a[TAB][TAB]

The bash shell would then output the candidate:
	add        am         annotate   apply      archive

If we then typed the following at the bash shell:
	git ap[TAB][TAB]

The bash shell would then fully autocomplete it to:
	git apply

Again, that is a powerful feature for the end user!
And creates a better user experience (UX) for the command line interface!

How Autocompletion Support Works

Some people add autocomplete support to their programs, they create a separate (often complex) shell script.
(Maybe a bash shell script. Although not necessarily)

This is NOT the way this package approaches this.

This package takes the approach of having everything self-contained.
And making it so the autocomplete support comes from the same program that is being autocompleted.

There are 3 parts to get this to work.
I.e., there are 3 tricks to this.

The first part of this trick is to use ‘complete’.

I.e., to put a command such as:

	complete -C '/the/full/path/to/myprogram compgen' myprogram

Into a file such as:

• ~/.bashrc

• ~/.bash_profile

• ~/.bash_login

• ~/.profile

• ~/.zshrc

• $XDG_CONFIG_HOME/fish/completions/

• ~/.config/fish/completions/

Where you put it can depend on the operating system (OS) you (or the end user) are on.
And also what command line shell you (or the end user) are using.

(Don't worry, this package provides tool to make that happen automagically.
This description is being included as documentation for those curious about how it works.)

But anyway, the second part of this trick is to make your program output data in a very specific format when it is called with the “compgen” command.

(Note, there is nothing special about the “compgen” command. I had to pick way of triggering the program to output information for autocomplete.
I chose to call it “compgen”.
I could have called it almost anything.
But (somewhat) arbitrarily called it “compgen”, because many bash scripts that do something like this make use of a program called “compgen”.
I thought it might be helpful to use the same name.)

The third part of the trick is for your program, when handling an autocomplete, to make use of these environment variables:

• COMP_KEY

• COMP_LINE

• COMP_POINT

• COMP_TYPE

And based on what is in those environment variables, change (and possibly narrow) what you output as candidates for the autocompletion.

(Again, don't worry, this package provides tool to make that happen automagically.
This description is being included as documentation for those curious about how it works.)

*/
package cliautocomplete
