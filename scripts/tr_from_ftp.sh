#!/bin/bash
# Handle fundamental files from Reuters ftp
# Fetches, parses, and submits to database 

START=$SECONDS

# allow quick quit
control_c()
{
	echo -e "\n\nProcess cancelled.. exiting"
	exit $?
}
trap control_c SIGINT

# allow user to select a date (YYYYMMDD)
if [ $# -gt 0 ]
then
	TODAY=$1
	RECENT=$2
else
	TODAY=`date +%Y%m%d`
	RECENT="--recent"
fi

HOST='rkd.knowledge.reuters.com'
USER='your_ftp_username_here'
PASS='your_ftp_password_here'

# tr includes a lot of files in the ftp
# use this line to filter out only the most recent archive
SEARCH=$TODAY"_1of1.xml.zip"

OUTPUT='ls.output'
ODIR=/home/data/reuters
CUR=`pwd`

# jump to data dir
cd $ODIR
#rm -vr $ODIR/

# get a list of available files
ftp -in $HOST <<END_SESSION
quote USER $USER
quote PASS $PASS
ls . $OUTPUT
quit
END_SESSION

# select only the files we care about 
for fname in `grep $SEARCH $OUTPUT | awk '{print $9;}'`
do
	SELECTED+=" $fname"
	echo $fname
done

# download the selected files
ftp -in $HOST <<END_SESSION
quote USER $USER
quote PASS $PASS
binary
mget $SELECTED
quit
END_SESSION

# unzip the files from today
for fname in `ls *$SEARCH`
do 
	UNZIPPED=`basename $fname .xml.zip`
	unzip -uo $fname -d $UNZIPPED

	# process the downloaded files
	find $UNZIPPED -name STDINT_* -exec /usr/bin/python /home/deploy/investigator/bin/parseReuters.py $RECENT --files {} +

done

# return to starting dir
cd $CUR

DURATION=$(($SECONDS - $START))
echo "Data fetched and ingested in $DURATION seconds"
