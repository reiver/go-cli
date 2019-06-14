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

(And not output any candidates, since it was able to do a full autocomplete.)

Autocomplete Flags

That type of autocomplete that we went over is automatic in command line shells that support autocomplete, such as the bash shell.

I.e., your program does not need to do anything to support that syle of autocomplete.

However, there is another kind of autocomplete that is very useful to the end user, but your program would have to do something extra
(such as integrate the tools provided by this package) to make it work.

But before going over that, let's describe what this other kind of autocomplete is like.

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


*/
package cliautocomplete
