package fr.oxodao.ledstrip;

import org.apache.http.NameValuePair;
import org.apache.http.client.HttpClient;
import org.apache.http.client.entity.UrlEncodedFormEntity;
import org.apache.http.client.methods.HttpPost;
import org.apache.http.impl.client.HttpClientBuilder;
import org.apache.http.message.BasicNameValuePair;

import java.io.IOException;
import java.util.ArrayList;
import java.util.Collections;
import java.util.List;

public class ApiCall {

    private void apiCall(String endpoint, BasicNameValuePair ...values)
    {
        String url = Ledstrip.getInstance().settings.url;
        HttpClient client = HttpClientBuilder.create().build();
        HttpPost post = new HttpPost(url + (url.endsWith("/") ? "" : "/") + "api/" + endpoint);
        try {
            List<NameValuePair> nvp = new ArrayList<>();
            Collections.addAll(nvp, values);
            post.setEntity(new UrlEncodedFormEntity(nvp));

            client.execute(post);
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    public void setColor(String color) {
        this.apiCall("color/set", new BasicNameValuePair("color", color));
    }

    public void setBrightness(int brightness) {
        this.apiCall("brightness/set", new BasicNameValuePair("brightness", ""+brightness));
    }
    public void spark(String color, int duration) {
        this.apiCall("color/spark", new BasicNameValuePair("color", color), new BasicNameValuePair("duration", ""+duration));
    }

}
