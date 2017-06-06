Imports System.Text
Imports System.Net
Imports System.IO

Module Startup

    Public HDFSServer As String
    Public HDFSPort As String
    Public BaseImportFolderName As String
    Public destFolderName As String

    Sub Main()


        HDFSServer = "http://tcambari04"
        HDFSPort = "50070"


        Dim RowHeaderNumber As Integer = 2
        Dim IngestData As New DataSet
        IngestData.DataSetName = "IngestData"

        For Each fi As String In System.IO.Directory.GetFiles("C:\Temp\EDI\", "*orig.txt")
            Dim CurrentTableName As String = ""

            Using MyReader As New Microsoft.VisualBasic.FileIO.TextFieldParser(fi)

                MyReader.TrimWhiteSpace = True
                MyReader.TextFieldType = FileIO.FieldType.Delimited
                MyReader.SetDelimiters(vbTab)
                Dim currentRow As String()

                While Not MyReader.EndOfData


                    Try

                        currentRow = MyReader.ReadFields()

                        If MyReader.LineNumber < RowHeaderNumber + 1 And currentRow.Length < 2 Then 'Always remember that each execution of readfields increments the line number even when using the previous line
                            'This is the title of the EDI file, we will use that for our datatable name
                            CurrentTableName = currentRow(0)
                            IngestData.Tables.Add(CurrentTableName)
                        ElseIf MyReader.LineNumber = RowHeaderNumber + 1 Then
                            For Each colname As String In currentRow
                                IngestData.Tables(CurrentTableName).Columns.Add(colname) ' add each column to the table in order
                            Next

                        ElseIf MyReader.LineNumber > RowHeaderNumber + 1 Then
                            Dim RollThroughColumns As Integer = 0
                            Dim dr As DataRow = IngestData.Tables(CurrentTableName).NewRow
                            For Each fieldvalue As String In currentRow
                                dr(RollThroughColumns) = fieldvalue
                                RollThroughColumns += 1
                            Next
                            IngestData.Tables(CurrentTableName).Rows.Add(dr)
                        End If

                    Catch ex As Microsoft.VisualBasic.
                                FileIO.MalformedLineException
                        Console.WriteLine("Line " & ex.Message &
                        "is not valid and will be skipped.")
                    Catch exn As Exception
                        Console.WriteLine("Unhandled Exception:" & exn.ToString)
                    End Try
                End While
            End Using



        Next

        IngestData.WriteXml("IngestData.xml")
        'Send data to hive
        Dim f As FileInfo = New System.IO.FileInfo("IngestData.xml")
        Dim FileSize As Long = f.Length
        UploadFile("EDIData/", "IngestData.xml", f, True)


    End Sub


    Public Sub UploadFile(ByVal DestinationFolder As String, ByVal DestinationFile As String, ByVal SourceFile As FileInfo, Optional ByVal Overwrite As Boolean = False)

        'First step, post to the HDFS 

        Dim FirstURL As String = HDFSServer & ":" & HDFSPort
        ' FirstURL += "/webhdfs/v1/" & destFolderName & "/?op=CREATE&overwrite=true&blocksize=" & BlockSize.ToString & "&buffersize=0"
        FirstURL += "/webhdfs/v1/" & DestinationFolder & SourceFile.Name & "?user.name=hdfs&op=CREATE&overwrite=TRUE"

        Dim requestGetStorageNode As HttpWebRequest = CType(WebRequest.Create(FirstURL), HttpWebRequest)
        requestGetStorageNode.Method = "PUT"
        requestGetStorageNode.AllowAutoRedirect = False
        requestGetStorageNode.GetRequestStream()
        Dim responseGetStorageNode As HttpWebResponse = CType(requestGetStorageNode.GetResponse(), HttpWebResponse)

        Dim request2 As HttpWebRequest = CType(WebRequest.Create(responseGetStorageNode.Headers("Location").ToString), HttpWebRequest)
        request2.Method = "PUT"
        Console.WriteLine(request2.RequestUri)
        If SourceFile.Length < 500000000 Then
            Try

                Using fileStream As FileStream = File.OpenRead(SourceFile.FullName)

                    Using requestStream As Stream = request2.GetRequestStream()
                        Dim bufferSize As Integer = fileStream.Length
                        Dim buffer(bufferSize - 1) As Byte
                        Dim byteCount As Integer = 0
                        byteCount = fileStream.Read(buffer, 0, bufferSize)
                        Do While byteCount > 0
                            requestStream.Write(buffer, 0, byteCount)
                            byteCount = fileStream.Read(buffer, 0, bufferSize)
                        Loop
                    End Using
                End Using

                Dim response2 As HttpWebResponse = CType(request2.GetResponse(), HttpWebResponse)

                Dim receiveStream As Stream = responseGetStorageNode.GetResponseStream()

                ' Pipes the stream to a higher level stream reader with the required encoding format. 
                Dim readStream As New StreamReader(receiveStream, Encoding.UTF8)

                '  Console.WriteLine("Response stream received.")
                Console.WriteLine(readStream.ReadToEnd())
                responseGetStorageNode.Close()
                readStream.Close()


            Catch ex As Exception
                Console.WriteLine("Skipped: " & ex.Message & SourceFile.FullName)

            End Try

            'Console.WriteLine("Content length is {0}", responseGetStorageNode.ContentLength)
            ' Console.WriteLine("Content type is {0}", responseGetStorageNode.ContentType)
        Else
            Console.WriteLine("Skipped big: " & SourceFile.FullName)
        End If

    End Sub



End Module
 
