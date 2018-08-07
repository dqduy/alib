package com.sevenlib.app.imageviewer;

import android.support.v7.widget.RecyclerView;
import android.view.View;
import android.view.ViewGroup;
import android.widget.TextView;
import android.widget.Button;
import android.view.LayoutInflater;

import java.util.ArrayList;
import java.util.List;

public class ImageItemsAdapter  extends RecyclerView.Adapter<ImageItemsAdapter.ViewHolder> {

	public class ViewHolder extends RecyclerView.ViewHolder {
		public TextView url;
		public Button name;

		public ViewHolder(View item) {
			super(item);

			url  = (TextView) item.findViewById(R.id.url);
			name = (Button) item.findViewById(R.id.name);
		}
	}

	private List<ImageItem> list;

	public ImageItemsAdapter(List<ImageItem> list) {
		this.list = list;
	}

	@Override
	public ImageItemsAdapter.ViewHolder onCreateViewHolder(ViewGroup parent, int viewType) {
		LayoutInflater inflater = LayoutInflater.from(parent.getContext());

		//Load item skeleton
		View imageView = inflater.inflate(R.layout.item, parent, false);

		ViewHolder viewHolder = new ViewHolder(imageView);
		return viewHolder;
	}

	@Override
	public void onBindViewHolder(ImageItemsAdapter.ViewHolder viewHolder, int position) {
		ImageItem item = list.get(position);

		TextView url = viewHolder.url;
		url.setText(item.url);
		Button name = viewHolder.name;
		name.setText(item.name);
	}

	@Override
	public int getItemCount() {
		return list.size();
	}
}