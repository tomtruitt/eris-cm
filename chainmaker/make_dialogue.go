package chains

func ChainsMakeWelcome() string {
	return `Welcome! I'm the marmot that helps you make your chain.

eris chains is your gateway to permissioned, smart contract compatible chains.
There is a bit of complexity around getting these chains set up. It is my
marmot-y task to make this as easy as we can.

First we will run through the eris chains typical account types and I'll ask
you how many of each account type you would like.

After that you will have an opportunity to create your own account types and
add those groupings into the genesis.json.

Remember, I'm only useful when you are making a new chain. After your chain
is established then you can modify the permissions using other eris tooling.

Are you ready to make your own chain (Y/n)? `
}

func ChainsMakePrelimQuestions() map[string]string {
	questions := map[string]string{
		"dryrun": "Would you like to review each of the account groups before making the chain? (y/N) ",
		"tokens": "Do you care about setting the number of tokens for each account type? If not, the marmots will set reasonable defaults for you. (y/N) ",
		"manual": "After the built in account groups would you like to make your manual own account groups with fine tuned permissions? (y/N) ",
	}
	return questions
}

func ChainsMakeRoot() string {
	return `The Root Group.

Group Definition:

Users who have a key which is registered with root privileges can do everything
on the chain. They have all of the permissions possible.

Who Should Get These?:

If you are making a small chain just to play around then usually you would
give all of the accounts needed for your experiment root privileges (unless you
were testing different) privilege types.

If you are making a more complex chain, then you would usually have a few
keys which have registered root permissions and as such will act in a capacity
similar to a network administrator in other data management situations.

How many keys do you want in the Root Group? (3) `
}

func ChainsMakeRootTokens() string {
	return "How many tokens should each key in the Root Group be given? "
}

func ChainsMakeDevelopers() string {
	return `The Developer Group.

Group Definition:

Users who have a key which is registered with developer privileges can send
tokens; call contracts; create contracts; create accounts; use the name registry;
and modify account's roles.

Who Should Get These?:

Generally the development team seeking to build the application on top of the
given chain would be within the group. If this is a multi organizational
chain then developers from each of the stakeholders should generally be registered
within this group, although that design is up to you.

How many keys do you want in the Developer Group? (6) `
}

func ChainsMakeDevelopersTokens() string {
	return "How many tokens should each key in the Developer Group be given? "
}

func ChainsMakeValidators() string {
	return `The Validators Group.

Group Definition:

Users who have a key which is registered with validator privileges can
only post a bond and begin validating the chain. This is the only privilege
this account group gets.

Who Should Get These?:

Generally the marmots recommend that you put your validator nodes onto a cloud
(IaaS) provider so that they will be always running.

We also recommend that if you are in a multi organizational chain then you would
have some separation of the validators to be ran by the different organizations
in the system.

How many keys do you want in the Validators Group? (7) `
}

func ChainsMakeValidatorsTokens() string {
	return "How many tokens should each key in the Validators Group be given? "
}

func ChainsMakeParticipants() string {
	return `The Participants Group.

Group Definition:

Users who have a key which is registered with participant privileges can send
tokes; call contracts; and use the name registry.

Who Should Get These?:

Generally the number of participants in your chain who do not need elevated
privileges should be given these keys.

Usually this group will have the most number of keys of all of the groups.

How many keys do you want in the Participants Group? (25) `
}

func ChainsMakeParticipantsTokens() string {
	return "How many tokens should each key in the Participants Group be given? "
}

func ChainsMakeManual() string {
	return `Make Your Own Group.

Group Definition:

You will next be asked a series of questions regarding this group as to
what permissions you want.

Who Should Get These?:

Don't ask us, you are the one that wanted "manual" :-)

For more on eris chains permissions see here:

https://docs.erisindustries.com/documentation/eris-db-permissions/

How many keys do you want in *this* manual group? (You can make more than one manual group) `
}

func ChainsMakeManualTokens() string {
	return "How many tokens should each key in this manual group be given? "
}

// todo: this should autopopulate, but for later.
func ChainsMakeManualQuestions() []string {
	return []string{
		"Does the group have root privileges? (y/n) ",
		"Does the group have send privileges? (y/n) ",
		"Does the group have call privileges? (y/n) ",
		"Does the group have create_contract privileges? (y/n) ",
		"Does the group have create_account privileges? (y/n) ",
		"Does the group have bond privileges? (y/n) ",
		"Does the group have name privileges? (y/n) ",
		"Does the group have has_base privileges? (y/n) ",
		"Does the group have set_base privileges? (y/n) ",
		"Does the group have unset_base privileges? (y/n) ",
		"Does the group have set_global privileges? (y/n) ",
		"Does the group have has_role privileges? (y/n) ",
		"Does the group have add_role privileges? (y/n) ",
		"Does the group have rm_role privileges? (y/n) ",
	}
}

func ChainsMakeManualMore() string {
	return "Do you want to make another manual group (y/n) "
}
