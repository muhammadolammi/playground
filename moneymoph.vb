Imports System.IO
Imports System.Net
Imports Newtonsoft.Json.Linq
Imports System.Text


Public Class MoneyMoph
    ' Delclare Countries data List
    Private astCountriesData(34, 1) As String
    Private stOtherCountries(4, 2) As String


    Private stFromCountry As String
    Private stFromCountryCurrency As String

    Private stToCountry As String
    Private stToCountryCurrency As String

    Private Function DesignPage()
        ' Lets Designs
        Panel1.BackColor = ColorTranslator.FromHtml("#FCFCFC")
        ConvertText.BackColor = ColorTranslator.FromHtml("#FCFCFC")

        LblHeaderName.ForeColor = ColorTranslator.FromHtml("#010055")
        ConvertLabel.ForeColor = ColorTranslator.FromHtml("#7D7D7D")
        ToLabel.ForeColor = ColorTranslator.FromHtml("#7D7D7D")
        BtnConvert.BackColor = ColorTranslator.FromHtml("#438FE7")

    End Function
    Private Sub Panel1_Paint(sender As Object, e As PaintEventArgs) Handles Panel1.Paint
        ControlPaint.DrawBorder(e.Graphics, Panel1.ClientRectangle, ColorTranslator.FromHtml("#7AC7FE"), ButtonBorderStyle.Solid)
    End Sub



    Private Sub MoneyMoph_Load(sender As Object, e As EventArgs) Handles MyBase.Load
        ' Disable the test button you can enable later
        BtnTest.Visible = False
        'Design The Page
        DesignPage()
        ' Set the screen to full screen on load
        Me.WindowState = FormWindowState.Maximized
        ' Set the countries Data
        astCountriesData(0, 0) = "Australia"
        astCountriesData(0, 1) = "AUD"

        astCountriesData(1, 0) = "Botswana"
        astCountriesData(1, 1) = "BWP"

        astCountriesData(2, 0) = "Brazil"
        astCountriesData(2, 1) = "BRL"

        astCountriesData(3, 0) = "Canada"
        astCountriesData(3, 1) = "CAD"

        astCountriesData(4, 0) = "Chile"
        astCountriesData(4, 1) = "CLP"

        astCountriesData(5, 0) = "China"
        astCountriesData(5, 1) = "CNY"

        astCountriesData(6, 0) = "Czech Republic"
        astCountriesData(6, 1) = "CZK"

        astCountriesData(7, 0) = "Denmark"
        astCountriesData(7, 1) = "DKK"

        astCountriesData(8, 0) = "Egypt"
        astCountriesData(8, 1) = "EGP"

        astCountriesData(9, 0) = "Ghana"
        astCountriesData(9, 1) = "GHS"

        astCountriesData(10, 0) = "Hong Kong"
        astCountriesData(10, 1) = "HKD"

        astCountriesData(11, 0) = "India"
        astCountriesData(11, 1) = "INR"

        astCountriesData(12, 0) = "Indonesia"
        astCountriesData(12, 1) = "IDR"

        astCountriesData(13, 0) = "Israel"
        astCountriesData(13, 1) = "ILS"

        astCountriesData(14, 0) = "Japan"
        astCountriesData(14, 1) = "JPY"

        astCountriesData(15, 0) = "Kenya"
        astCountriesData(15, 1) = "KES"

        astCountriesData(16, 0) = "Kuwait"
        astCountriesData(16, 1) = "KWD"

        astCountriesData(17, 0) = "Malaysia"
        astCountriesData(17, 1) = "MYR"

        astCountriesData(18, 0) = "Mexico"
        astCountriesData(18, 1) = "MXN"

        astCountriesData(19, 0) = "Morocco"
        astCountriesData(19, 1) = "MAD"

        astCountriesData(20, 0) = "Nigeria"
        astCountriesData(20, 1) = "NGN"

        astCountriesData(21, 0) = "Oman"
        astCountriesData(21, 1) = "OMR"

        astCountriesData(22, 0) = "Philippines"
        astCountriesData(22, 1) = "PHP"

        astCountriesData(23, 0) = "Russia"
        astCountriesData(23, 1) = "RUB"

        astCountriesData(24, 0) = "Saudi Arabia"
        astCountriesData(24, 1) = "SAR"

        astCountriesData(25, 0) = "Singapore"
        astCountriesData(25, 1) = "SGD"

        astCountriesData(26, 0) = "South Africa"
        astCountriesData(26, 1) = "ZAR"

        astCountriesData(27, 0) = "South Korea"
        astCountriesData(27, 1) = "KRW"

        astCountriesData(28, 0) = "Switzerland"
        astCountriesData(28, 1) = "CHF"

        astCountriesData(29, 0) = "Thailand"
        astCountriesData(29, 1) = "THB"

        astCountriesData(30, 0) = "Tunisia"
        astCountriesData(30, 1) = "TND"

        astCountriesData(31, 0) = "Turkey"
        astCountriesData(31, 1) = "TRY"

        astCountriesData(32, 0) = "United Arab Emirates"
        astCountriesData(32, 1) = "AED"

        astCountriesData(33, 0) = "United Kingdom"
        astCountriesData(33, 1) = "GBP"

        astCountriesData(34, 0) = "United States"
        astCountriesData(34, 1) = "USD"
        ' Lets set the Combo box data programatically from the country list
        ' Clear existing items in the ComboBox
        ComboBoxFromCountry.Items.Clear()
        ' Loop through the astCountriesData array and add each country to the ComboBox
        For i As Integer = 0 To 34
            ComboBoxFromCountry.Items.Add(astCountriesData(i, 1))
        Next
        ' Clear existing items in the ComboBox
        ComboBoxToCountry.Items.Clear()
        ' Loop through the astCountriesData array and add each country to the ComboBox
        For i As Integer = 0 To 34
            ComboBoxToCountry.Items.Add(astCountriesData(i, 1))
        Next

        'SET Default SELECTED COUNTRIES Data To Nigeria and united state
        ComboBoxFromCountry.SelectedItem = astCountriesData(34, 1)
        ComboBoxToCountry.SelectedItem = astCountriesData(20, 1)

        ' Update the countries data initially until user update it 

        stFromCountry = astCountriesData(ComboBoxFromCountry.SelectedIndex, 0)
        stFromCountryCurrency = astCountriesData(ComboBoxFromCountry.SelectedIndex, 1)

        stToCountry = astCountriesData(ComboBoxToCountry.SelectedIndex, 0)
        stToCountryCurrency = astCountriesData(ComboBoxToCountry.SelectedIndex, 1)
    End Sub


    Private Function GetRandomCountiesIndexes(numOfIndexes As Integer) As List(Of Integer)
        Dim randomIndexes As New HashSet(Of Integer)
        Dim random As New Random()
        Dim fromIndex As Integer = ComboBoxFromCountry.SelectedIndex
        Dim toIndex As Integer = ComboBoxToCountry.SelectedIndex
        For i = 0 To numOfIndexes
            Dim randomIndex As Integer
            'Let check if the random number is already choosen index
            Do
                randomIndex = random.Next(0, 34) ' Adjust maxIndex to your maximum index range
            Loop While randomIndex = fromIndex Or randomIndex = toIndex Or randomIndexes.Contains(randomIndex)
            ' The do while block means that
            ' Set the random index to the number generated
            ' Kepp looping until the random index is non of from and to country indexes and also not already in the random indexes set
            ' This way we wont generate 2 same random index, and we wont deplicate chosen countries
            randomIndexes.Add(randomIndex)
        Next
        Return randomIndexes.ToList()
    End Function

    Private Function FetchData() As String
        Try
            ' SET OTHER countries randomly 
            Dim randomCountriesIndexes As List(Of Integer) = GetRandomCountiesIndexes(5)
            ' lETS Populate other countries
            For i = 0 To 4
                Dim iIndex As Integer = randomCountriesIndexes(i)
                stOtherCountries(i, 0) = astCountriesData(iIndex, 0)
                stOtherCountries(i, 1) = astCountriesData(iIndex, 1)
            Next
            Dim stCurrencies As New StringBuilder()
            stCurrencies.Append(stToCountryCurrency & ",")
            For i As Integer = 0 To stOtherCountries.GetLength(0) - 1
                Dim currencyString As String = stOtherCountries(i, 1).ToString().ToUpper() & ","
                stCurrencies.Append(currencyString)
            Next
            ' Remove the final comma and space
            If stCurrencies.Length > 0 Then
                stCurrencies.Length -= 1
            End If


            Dim stWebRequestCurrencies As String = stCurrencies.ToString()
            Dim url As String = ("https://api.currencyapi.com/v3/latest?base_currency=" & stFromCountryCurrency & "&currencies=" & stWebRequestCurrencies)
            Dim webRequest As WebRequest = WebRequest.Create(url)
            Dim request As HttpWebRequest = CType(webRequest, HttpWebRequest)
            request.Method = "GET"
            request.Headers.Add("apikey", "cur_live_e5g0kE4kW9fXWTkBjhpUQxT1RMullT5JO4qUySXX")
            request.ContentType = "application/json"
            Dim responseText As String
            Dim response As HttpWebResponse = CType(request.GetResponse(), HttpWebResponse)
            ' Read the response (assuming it's a string)
            Dim responseStream As Stream = response.GetResponseStream()
            Dim reader As New StreamReader(responseStream)
            responseText = reader.ReadToEnd()
            ' Close the reader and the response
            reader.Close()
            response.Close()
            Return responseText
        Catch ex As WebException
            ' Handle web exceptions
            Dim errorResponse As HttpWebResponse = CType(ex.Response, HttpWebResponse)
            If errorResponse IsNot Nothing Then
                MsgBox("Error: " & errorResponse.StatusCode & " - " & errorResponse.StatusDescription)
                Return ""
            Else
                MsgBox("Error: " & ex.Message)
                Return ""
            End If
        Catch ex As Exception
            ' Handle other exceptions 
            MsgBox("An unexpected error occurred: " & ex.Message)
            Return ""

        End Try
    End Function


    Private Sub BtnConvert_Click(sender As Object, e As EventArgs) Handles BtnConvert.Click
        Dim iAmountToConvert As Integer
        ' Send Api request to get data
        Dim result As Decimal
        ' Lets start getting the data and handle wrong data from amount
        If IsNumeric(ConvertText.Text) = True Then
            iAmountToConvert = CInt(ConvertText.Text)
        Else
            MsgBox("Your amount must be a number")
            Exit Sub
        End If
        ' Fetch currency data as JSON string
        Dim data As String = FetchData()
        If String.IsNullOrEmpty(data) Then
            MsgBox("We can't get data from backend, try again.")
            Exit Sub
        End If
        ' Parse the JSON string
        Dim json As JObject = JObject.Parse(data)
        Dim currencies As JObject = json("data")
        If currencies IsNot Nothing AndAlso currencies.Count = 6 Then
            result = currencies(stToCountryCurrency).Value(Of Decimal)("value")
            For i As Integer = 0 To stOtherCountries.GetLength(0) - 1
                    Dim countryCurrencyvalue = currencies(stOtherCountries(i, 1)).Value(Of Decimal)("value")
                    stOtherCountries(i, 2) = countryCurrencyvalue
                Next
           
        Else
            MsgBox("Currencies is not complete.")
            Exit Sub
        End If

        ' Calculate converted amount
        Dim decConvertedAmount = iAmountToConvert * result
        If decConvertedAmount > 1 Then
            decConvertedAmount = Math.Round(decConvertedAmount)
        Else
            decConvertedAmount = Math.Round(decConvertedAmount, 4)
        End If

        If result > 1 Then
            result = Math.Round(result)
        Else
            result = Math.Round(result, 4)
        End If
        ' Update UI with results
        ConvertedToText.Text = decConvertedAmount
        CDFromCurrencyName.Text = stFromCountryCurrency
        CDFromCurrencyCountry.Text = stFromCountry
        CDFromCurrencyValue.Text = $"1 {stFromCountryCurrency}"
        CDToCurrencyName.Text = stToCountryCurrency
        CDToCurrencyCountry.Text = stToCountry
        CDToCurrencyValue.Text = $"{result } {stToCountryCurrency}" 

        CDOtherCountry0CurrencyValue.Text = $"{Math.Round(stOtherCountries(0, 2) * iAmountToConvert)}"
        CDOtherCountry0CurrencyCountry.Text = $"{stOtherCountries(0, 0)}"
        CDOtherCountry0CurrencyName.Text = $"{stOtherCountries(0, 1)}"


        CDOtherCountry1CurrencyValue.Text = $"{Math.Round(stOtherCountries(1, 2) * iAmountToConvert)}"
        CDOtherCountry1CurrencyCountry.Text = $"{stOtherCountries(1, 0)}"
        CDOtherCountry1CurrencyName.Text = $"{stOtherCountries(1, 1)}"

        CDOtherCountry2CurrencyValue.Text = $"{Math.Round(stOtherCountries(2, 2) * iAmountToConvert)}"
        CDOtherCountry2CurrencyCountry.Text = $"{stOtherCountries(2, 0)}"
        CDOtherCountry2CurrencyName.Text = $"{stOtherCountries(2, 1)}"

        CDOtherCountry3CurrencyValue.Text = $"{Math.Round(stOtherCountries(3, 2) * iAmountToConvert)}"
        CDOtherCountry3CurrencyCountry.Text = $"{stOtherCountries(3, 0)}"
        CDOtherCountry3CurrencyName.Text = $"{stOtherCountries(3, 1)}"

        CDOtherCountry4CurrencyValue.Text = $"{Math.Round(stOtherCountries(4, 2) * iAmountToConvert)}"
        CDOtherCountry4CurrencyCountry.Text = $"{stOtherCountries(4, 0)}"
        CDOtherCountry4CurrencyName.Text = $"{stOtherCountries(4, 1)}"
    End Sub



  


    'Update Selected Country when they are changed
    Private Sub ComboBoxFromCountry_SelectedIndexChanged(sender As Object, e As EventArgs) Handles ComboBoxFromCountry.SelectedIndexChanged
        ' SET country data 
        stFromCountry = astCountriesData(ComboBoxFromCountry.SelectedIndex, 0)
        stFromCountryCurrency = astCountriesData(ComboBoxFromCountry.SelectedIndex, 1)
    End Sub

    Private Sub ComboBoxToCountry_SelectedIndexChanged(sender As Object, e As EventArgs) Handles ComboBoxToCountry.SelectedIndexChanged
        ' SET country data
        stToCountry = astCountriesData(ComboBoxToCountry.SelectedIndex, 0)
        stToCountryCurrency = astCountriesData(ComboBoxToCountry.SelectedIndex, 1)
    End Sub



    Private Sub BtnTest_Click(sender As Object, e As EventArgs) Handles BtnTest.Click


        Dim aiRandomIndexes As List(Of Integer) = GetRandomCountiesIndexes(5)
        Dim stReturnText As New StringBuilder()
        For Each index In aiRandomIndexes
            stReturnText.Append(index.ToString()).Append(" ")



        Next
        MsgBox(stReturnText.ToString())

    End Sub

    Private Sub BtnAvailableCountries_Click(sender As Object, e As EventArgs) Handles BtnAvailableCountries.Click

        Dim stReturnString As String = ""
        ' Loop through the astCountriesData array and add each country to the ComboBox
        For i As Integer = 0 To 34
            stReturnString = $"{stReturnString} {astCountriesData(i, 0)}({astCountriesData(i, 1)}), "
        Next

        MsgBox(stReturnString)

    End Sub

    Private Sub BtnEndApp_Click(sender As Object, e As EventArgs) Handles BtnEndApp.Click
        End
    End Sub
End Class
