# gov_notify_go
send sms to coma separated list via gov notify (uk gov service)

to use call the exe with the followin argumens
-key : the API key string that you have configured in Gov notify, use the full key in format {key_name}-{iss-uuid}-{secret-key-uuid}

-phoneNumbers : comma separated list of phone numbers that SMS is to be sent to

-messageID : the GUID of the message template to use

-messageContent : message personaliasation json block (remember escape chars where needed), eg: "{\\""first_name\\"": \\""Amala\\"",\\""appointment_date\\"": \\""1 January 2018 at 1:00PM\\"",}"